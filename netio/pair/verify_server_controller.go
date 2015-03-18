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
func (verify *VerifyServerController) SharedKey() [32]byte {
	return verify.session.SharedKey
}

func (verify *VerifyServerController) KeyVerified() bool {
	return verify.step == VerifyStepFinishResponse
}

func (verify *VerifyServerController) Handle(in common.Container) (common.Container, error) {
	var out common.Container
	var err error

	method := PairMethodType(in.GetByte(TagPairingMethod))

	// It is valid that method is not sent
	// If method is sent then it must be 0x00
	if method != PairingMethodDefault {
		return nil, ErrInvalidPairMethod(method)
	}

	seq := VerifyStepType(in.GetByte(TagSequence))

	switch seq {
	case VerifyStepStartRequest:
		if verify.step != VerifyStepWaiting {
			verify.Reset()
			return nil, ErrInvalidInternalVerifyStep(verify.step)
		}
		out, err = verify.handlePairVerifyStart(in)
	case VerifyStepFinishRequest:
		if verify.step != VerifyStepStartResponse {
			verify.Reset()
			return nil, ErrInvalidInternalVerifyStep(verify.step)
		}

		out, err = verify.handlePairVerifyFinish(in)
	default:
		return nil, ErrInvalidVerifyStep(seq)
	}

	return out, err
}

// Client -> Server
// - Public key `A`
//
// Server
// - Create public `B` and secret key `S` based on `A`

// Server -> Client
// - B: server public key
// - signature: from server session public key, server name, client session public key
func (verify *VerifyServerController) handlePairVerifyStart(in common.Container) (common.Container, error) {
	verify.step = VerifyStepStartResponse

	clientPublicKey := in.GetBytes(TagPublicKey)
	log.Println("[VERB] ->     A:", hex.EncodeToString(clientPublicKey))
	if len(clientPublicKey) != 32 {
		return nil, ErrInvalidClientKeyLength
	}

	var otherPublicKey [32]byte
	copy(otherPublicKey[:], clientPublicKey)

	verify.session.GenerateSharedKeyWithOtherPublicKey(otherPublicKey)
	verify.session.SetupEncryptionKey([]byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))

	bridge := verify.context.GetBridge()
	LTSK := bridge.PrivateKey()

	material := make([]byte, 0)
	material = append(material, verify.session.PublicKey[:]...)
	material = append(material, bridge.Id()...)
	material = append(material, clientPublicKey...)
	signature, err := crypto.ED25519Signature(LTSK, material)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Encrypt
	encryptedOut := common.NewTLV8Container()
	encryptedOut.SetString(TagUsername, bridge.Id())
	encryptedOut.SetBytes(TagSignature, signature)

	encryptedBytes, mac, _ := crypto.Chacha20EncryptAndPoly1305Seal(verify.session.EncryptionKey[:], []byte("PV-Msg02"), encryptedOut.BytesBuffer().Bytes(), nil)

	out := common.NewTLV8Container()
	out.SetByte(TagSequence, verify.step.Byte())
	out.SetBytes(TagPublicKey, verify.session.PublicKey[:])
	out.SetBytes(TagEncryptedData, append(encryptedBytes, mac[:]...))

	log.Println("[VERB]        K:", hex.EncodeToString(verify.session.EncryptionKey[:]))
	log.Println("[VERB]        B:", hex.EncodeToString(verify.session.PublicKey[:]))
	log.Println("[VERB]        S:", hex.EncodeToString(verify.session.PrivateKey[:]))
	log.Println("[VERB]   Shared:", hex.EncodeToString(verify.session.SharedKey[:]))

	log.Println("[VERB] <-     B:", hex.EncodeToString(out.GetBytes(TagPublicKey)))

	return out, nil
}

// Client -> Server
// - Encrypted tlv8: username and signature
//
// Server enrypty tlv8 and validates signature

// Server -> Client
// - only sequence number
// - error code (optional)
func (verify *VerifyServerController) handlePairVerifyFinish(in common.Container) (common.Container, error) {
	verify.step = VerifyStepFinishResponse

	data := in.GetBytes(TagEncryptedData)
	message := data[:(len(data) - 16)]
	var mac [16]byte
	copy(mac[:], data[len(message):]) // 16 byte (MAC)
	log.Println("[VERB] ->     Message:", hex.EncodeToString(message))
	log.Println("[VERB] ->     MAC:", hex.EncodeToString(mac[:]))

	decryptedBytes, err := crypto.Chacha20DecryptAndPoly1305Verify(verify.session.EncryptionKey[:], []byte("PV-Msg03"), message, mac, nil)

	out := common.NewTLV8Container()
	out.SetByte(TagSequence, verify.step.Byte())

	if err != nil {
		verify.Reset()
		log.Println("[ERRO]", err)
		out.SetByte(TagErrCode, ErrCodeAuthenticationFailed.Byte()) // return error 2
	} else {
		in, err := common.NewTLV8ContainerFromReader(bytes.NewBuffer(decryptedBytes))
		if err != nil {
			return nil, err
		}

		username := in.GetString(TagUsername)
		signature := in.GetBytes(TagSignature)
		log.Println("[VERB]     client:", username)
		log.Println("[VERB]  signature:", hex.EncodeToString(signature))

		entity := verify.database.EntityWithName(username)
		if entity == nil {
			return nil, common.NewErrorf("Client %s is unknown", username)
		}

		if len(entity.PublicKey()) == 0 {
			return nil, common.NewErrorf("No LTPK available for client %s", username)
		}

		material := make([]byte, 0)
		material = append(material, verify.session.OtherPublicKey[:]...)
		material = append(material, []byte(username)...)
		material = append(material, verify.session.PublicKey[:]...)

		if crypto.ValidateED25519Signature(entity.PublicKey(), material, signature) == false {
			log.Println("[WARN] signature is invalid")
			verify.Reset()
			out.SetByte(TagErrCode, ErrCodeUnknownPeer.Byte()) // return error 4
		} else {
			log.Println("[VERB] signature is valid")
		}
	}

	return out, nil
}

func (verify *VerifyServerController) Reset() {
	verify.step = VerifyStepWaiting
}
