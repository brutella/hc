package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"
	"github.com/brutella/log"

	"bytes"
	"encoding/hex"
)

// VerifyServerController verifies the stored client public key and negotiates a shared secret
// which is used encrypt the upcoming session.
//
// Verification fails when the client is not known, the public key for the client was not found,
// or the packet's seal could not be verified.
type VerifyServerController struct {
	database db.Database
	context  netio.HAPContext
	session  *VerifySession
	step     VerifyStepType
}

func NewVerifyServerController(database db.Database, context netio.HAPContext) *VerifyServerController {
	controller := VerifyServerController{
		database: database,
		context:  context,
		session:  NewVerifySession(),
		step:     VerifyStepWaiting,
	}

	return &controller
}
func (c *VerifyServerController) SharedKey() [32]byte {
	return c.session.SharedKey
}

func (c *VerifyServerController) KeyVerified() bool {
	return c.step == VerifyStepFinishResponse
}

func (c *VerifyServerController) Handle(cont_in common.Container) (common.Container, error) {
	var cont_out common.Container
	var err error

	method := PairMethodType(cont_in.GetByte(TagPairingMethod))

	// It is valid that method is not sent
	// If method is sent then it must be 0x00
	if method != PairingMethodDefault {
		return nil, ErrInvalidPairMethod(method)
	}

	seq := VerifyStepType(cont_in.GetByte(TagSequence))

	switch seq {
	case VerifyStepStartRequest:
		if c.step != VerifyStepWaiting {
			c.Reset()
			return nil, ErrInvalidInternalVerifyStep(c.step)
		}
		cont_out, err = c.handlePairVerifyStart(cont_in)
	case VerifyStepFinishRequest:
		if c.step != VerifyStepStartResponse {
			c.Reset()
			return nil, ErrInvalidInternalVerifyStep(c.step)
		}

		cont_out, err = c.handlePairVerifyFinish(cont_in)
	default:
		return nil, ErrInvalidVerifyStep(seq)
	}

	return cont_out, err
}

// Client -> Server
// - Public key `A`
//
// Server
// - Create public `B` and secret key `S` based on `A`

// Server -> Client
// - B: server public key
// - signature: from server session public key, server name, client session public key
func (c *VerifyServerController) handlePairVerifyStart(cont_in common.Container) (common.Container, error) {
	c.step = VerifyStepStartResponse

	clientPublicKey := cont_in.GetBytes(TagPublicKey)
	log.Println("[VERB] ->     A:", hex.EncodeToString(clientPublicKey))
	if len(clientPublicKey) != 32 {
		return nil, ErrInvalidClientKeyLength
	}

	var otherPublicKey [32]byte
	copy(otherPublicKey[:], clientPublicKey)

	c.session.GenerateSharedKeyWithOtherPublicKey(otherPublicKey)
	c.session.SetupEncryptionKey([]byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))

	bridge := c.context.GetBridge()
	LTSK := bridge.SecretKey

	material := make([]byte, 0)
	material = append(material, c.session.PublicKey[:]...)
	material = append(material, bridge.Id()...)
	material = append(material, clientPublicKey...)
	signature, _ := crypto.ED25519Signature(LTSK, material)

	// Encrypt
	tlv_encrypt := common.NewTLV8Container()
	tlv_encrypt.SetString(TagUsername, bridge.Id())
	tlv_encrypt.SetBytes(TagSignature, signature)

	encrypted, mac, _ := crypto.Chacha20EncryptAndPoly1305Seal(c.session.EncryptionKey[:], []byte("PV-Msg02"), tlv_encrypt.BytesBuffer().Bytes(), nil)

	cont_out := common.NewTLV8Container()
	cont_out.SetByte(TagSequence, c.step.Byte())
	cont_out.SetBytes(TagPublicKey, c.session.PublicKey[:])
	cont_out.SetBytes(TagEncryptedData, append(encrypted, mac[:]...))

	log.Println("[VERB]       K:", hex.EncodeToString(c.session.EncryptionKey[:]))
	log.Println("[VERB]        B:", hex.EncodeToString(c.session.PublicKey[:]))
	log.Println("[VERB]        S:", hex.EncodeToString(c.session.SecretKey[:]))
	log.Println("[VERB]   Shared:", hex.EncodeToString(c.session.SharedKey[:]))

	log.Println("[VERB] <-     B:", hex.EncodeToString(cont_out.GetBytes(TagPublicKey)))

	return cont_out, nil
}

// Client -> Server
// - Encrypted tlv8: username and signature
//
// Server enrypty tlv8 and validates signature

// Server -> Client
// - only sequence number
// - error code (optional)
func (c *VerifyServerController) handlePairVerifyFinish(cont_in common.Container) (common.Container, error) {
	c.step = VerifyStepFinishResponse

	data := cont_in.GetBytes(TagEncryptedData)
	message := data[:(len(data) - 16)]
	var mac [16]byte
	copy(mac[:], data[len(message):]) // 16 byte (MAC)
	log.Println("[VERB] ->     Message:", hex.EncodeToString(message))
	log.Println("[VERB] ->     MAC:", hex.EncodeToString(mac[:]))

	decrypted, err := crypto.Chacha20DecryptAndPoly1305Verify(c.session.EncryptionKey[:], []byte("PV-Msg03"), message, mac, nil)

	cont_out := common.NewTLV8Container()
	cont_out.SetByte(TagSequence, c.step.Byte())

	if err != nil {
		c.Reset()
		log.Println("[ERRO]", err)
		cont_out.SetByte(TagErrCode, ErrCodeAuthenticationFailed.Byte()) // return error 2
	} else {
		decrypted_buffer := bytes.NewBuffer(decrypted)
		cont_in, err := common.NewTLV8ContainerFromReader(decrypted_buffer)
		if err != nil {
			return nil, err
		}

		username := cont_in.GetString(TagUsername)
		signature := cont_in.GetBytes(TagSignature)
		log.Println("[VERB]     client:", username)
		log.Println("[VERB]  signature:", hex.EncodeToString(signature))

		client := c.database.EntityWithName(username)
		if client == nil {
			return nil, common.NewErrorf("Client %s is unknown", username)
		}

		if len(client.PublicKey()) == 0 {
			return nil, common.NewErrorf("No LTPK available for client %s", username)
		}

		material := make([]byte, 0)
		material = append(material, c.session.OtherPublicKey[:]...)
		material = append(material, []byte(username)...)
		material = append(material, c.session.PublicKey[:]...)

		if crypto.ValidateED25519Signature(client.PublicKey(), material, signature) == false {
			log.Println("[WARN] signature is invalid")
			c.Reset()
			cont_out.SetByte(TagErrCode, ErrCodeUnknownPeer.Byte()) // return error 4
		} else {
			log.Println("[VERB] signature is valid")
		}
	}

	return cont_out, nil
}

func (c *VerifyServerController) Reset() {
	c.step = VerifyStepWaiting
}
