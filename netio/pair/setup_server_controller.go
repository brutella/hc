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

// SetupServerController handles pairing with a client using SRP.
// The client has to known the bridge password to successfully pair.
// When pairigin was successful, the client's public key (refered as LTPK - long term public key)
// is stored in the database for later use.
//
// Pairing may fail because the password is wrong or the key exchange failed (e.g. packet seals or SRP key authenticator is wrong, ...).
type SetupServerController struct {
	bridge   *netio.Bridge
	session  *PairSetupServerSession
	step     PairStepType
	database db.Database
}

// NewSetupServerController returns a new pair setup controller.
func NewSetupServerController(bridge *netio.Bridge, database db.Database) (*SetupServerController, error) {

	session, err := NewPairSetupServerSession(bridge.Id(), bridge.Password())
	if err != nil {
		return nil, err
	}

	controller := SetupServerController{
		bridge:   bridge,
		session:  session,
		database: database,
		step:     PairStepWaiting,
	}

	return &controller, nil
}

func (c *SetupServerController) Handle(cont_in common.Container) (cont_out common.Container, err error) {
	method := PairMethodType(cont_in.GetByte(TagPairingMethod))

	// It is valid that method is not sent
	// If method is sent then it must be 0x00
	if method != PairingMethodDefault {
		return nil, ErrInvalidPairMethod(method)
	}

	seq := PairStepType(cont_in.GetByte(TagSequence))

	switch seq {
	case PairStepStartRequest:
		if c.step != PairStepWaiting {
			c.reset()
			return nil, ErrInvalidInternalPairStep(c.step)
		}

		cont_out, err = c.handlePairStart(cont_in)
	case PairStepVerifyRequest:
		if c.step != PairStepStartResponse {
			c.reset()
			return nil, ErrInvalidInternalPairStep(c.step)
		}

		cont_out, err = c.handlePairVerify(cont_in)
	case PairStepKeyExchangeRequest:
		if c.step != PairStepVerifyResponse {
			c.reset()
			return nil, ErrInvalidInternalPairStep(c.step)
		}

		cont_out, err = c.handleKeyExchange(cont_in)
	default:
		return nil, ErrInvalidPairStep(seq)
	}

	return cont_out, err
}

// Client -> Server
// - Auth start
//
// Server -> Client
// - B: server public key
// - s: salt
func (c *SetupServerController) handlePairStart(cont_in common.Container) (common.Container, error) {
	cont_out := common.NewTLV8Container()
	c.step = PairStepStartResponse

	cont_out.SetByte(TagSequence, c.step.Byte())
	cont_out.SetBytes(TagPublicKey, c.session.PublicKey)
	cont_out.SetBytes(TagSalt, c.session.Salt)

	log.Println("[VERB] <-     B:", hex.EncodeToString(cont_out.GetBytes(TagPublicKey)))
	log.Println("[VERB] <-     s:", hex.EncodeToString(cont_out.GetBytes(TagSalt)))

	return cont_out, nil
}

// Client -> Server
// - A: client public key
// - M1: proof
//
// Server -> client
// - M2: proof
// or
// - auth error
func (c *SetupServerController) handlePairVerify(cont_in common.Container) (common.Container, error) {
	cont_out := common.NewTLV8Container()
	c.step = PairStepVerifyResponse

	cont_out.SetByte(TagSequence, c.step.Byte())

	cpublicKey := cont_in.GetBytes(TagPublicKey)
	log.Println("[VERB] ->     A:", hex.EncodeToString(cpublicKey))

	err := c.session.SetupSecretKeyFromClientPublicKey(cpublicKey)
	if err != nil {
		return nil, err
	}

	cproof := cont_in.GetBytes(TagProof)
	log.Println("[VERB] ->     M1:", hex.EncodeToString(cproof))

	sproof, err := c.session.ProofFromClientProof(cproof)
	if err != nil || len(sproof) == 0 { // proof `M1` is wrong
		log.Println("[WARN] Proof M1 is wrong")
		c.reset()
		cont_out.SetByte(TagErrCode, ErrCodeAuthenticationFailed.Byte()) // return error 2
	} else {
		log.Println("[INFO] Proof M1 is valid")
		err := c.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
		if err != nil {
			return nil, err
		}

		// Return proof `M1`
		cont_out.SetBytes(TagProof, sproof)
	}

	log.Println("[VERB] <-     M2:", hex.EncodeToString(cont_out.GetBytes(TagProof)))
	log.Println("[VERB]         S:", hex.EncodeToString(c.session.SecretKey))
	log.Println("[VERB]         K:", hex.EncodeToString(c.session.EncryptionKey[:]))

	return cont_out, nil
}

// Client -> Server
// - encrypted tlv8: client LTPK, client name and signature (of H, client name, LTPK)
// - auth tag (mac)
//
// Server
// - Validate signature of encrpyted tlv8
// - Read and store client LTPK and name
//
// Server -> Client
// - encrpyted tlv8: bridge LTPK, bridge name, signature (of H2, bridge name, LTPK)
func (c *SetupServerController) handleKeyExchange(cont_in common.Container) (common.Container, error) {
	cont_out := common.NewTLV8Container()

	c.step = PairStepKeyExchangeResponse

	cont_out.SetByte(TagSequence, c.step.Byte())

	data := cont_in.GetBytes(TagEncryptedData)
	message := data[:(len(data) - 16)]
	var mac [16]byte
	copy(mac[:], data[len(message):]) // 16 byte (MAC)
	log.Println("[VERB] ->     Message:", hex.EncodeToString(message))
	log.Println("[VERB] ->     MAC:", hex.EncodeToString(mac[:]))

	decrypted, err := crypto.Chacha20DecryptAndPoly1305Verify(c.session.EncryptionKey[:], []byte("PS-Msg05"), message, mac, nil)

	if err != nil {
		c.reset()
		log.Println("[ERRO]", err)
		cont_out.SetByte(TagErrCode, ErrCodeUnknown.Byte()) // return error 1
	} else {
		decrypted_buffer := bytes.NewBuffer(decrypted)
		cont_in, err := common.NewTLV8ContainerFromReader(decrypted_buffer)
		if err != nil {
			return nil, err
		}

		username := cont_in.GetString(TagUsername)
		ltpk := cont_in.GetBytes(TagPublicKey)
		signature := cont_in.GetBytes(TagSignature)
		log.Println("[VERB] ->     Username:", username)
		log.Println("[VERB] ->     LTPK:", hex.EncodeToString(ltpk))
		log.Println("[VERB] ->     Signature:", hex.EncodeToString(signature))

		// Calculate `H`
		H, _ := crypto.HKDF_SHA512(c.session.SecretKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
		material := make([]byte, 0)
		material = append(material, H[:]...)
		material = append(material, []byte(username)...)
		material = append(material, ltpk...)

		if crypto.ValidateED25519Signature(ltpk, material, signature) == false {
			log.Println("[WARN] ed25519 signature is invalid")
			c.reset()
			cont_out.SetByte(TagErrCode, ErrCodeAuthenticationFailed.Byte()) // return error 2
		} else {
			log.Println("[VERB] ed25519 signature is valid")
			// Store client LTPK and name
			client := db.NewClient(username, ltpk)
			c.database.SaveClient(client)
			log.Printf("[INFO] Stored LTPK '%s' for client '%s'\n", hex.EncodeToString(ltpk), username)

			LTPK := c.bridge.PublicKey
			LTSK := c.bridge.SecretKey

			// Send username, LTPK, signature as encrypted message
			H2, err := crypto.HKDF_SHA512(c.session.SecretKey, []byte("Pair-Setup-Accessory-Sign-Salt"), []byte("Pair-Setup-Accessory-Sign-Info"))
			material = make([]byte, 0)
			material = append(material, H2[:]...)
			material = append(material, []byte(c.session.Username)...)
			material = append(material, LTPK...)

			signature, err := crypto.ED25519Signature(LTSK, material)
			if err != nil {
				return nil, err
			}

			tlvPairKeyExchange := common.NewTLV8Container()
			tlvPairKeyExchange.SetBytes(TagUsername, c.session.Username)
			tlvPairKeyExchange.SetBytes(TagPublicKey, LTPK)
			tlvPairKeyExchange.SetBytes(TagSignature, []byte(signature))

			log.Println("[VERB] <-     Username:", tlvPairKeyExchange.GetString(TagUsername))
			log.Println("[VERB] <-     LTPK:", hex.EncodeToString(tlvPairKeyExchange.GetBytes(TagPublicKey)))
			log.Println("[VERB] <-     Signature:", hex.EncodeToString(tlvPairKeyExchange.GetBytes(TagSignature)))

			encrypted, mac, _ := crypto.Chacha20EncryptAndPoly1305Seal(c.session.EncryptionKey[:], []byte("PS-Msg06"), tlvPairKeyExchange.BytesBuffer().Bytes(), nil)
			cont_out.SetByte(TagPairingMethod, 0)
			cont_out.SetByte(TagSequence, PairStepKeyExchangeRequest.Byte())
			cont_out.SetBytes(TagEncryptedData, append(encrypted, mac[:]...))

			c.reset()
		}
	}

	return cont_out, nil
}

func (c *SetupServerController) reset() {
	c.step = PairStepWaiting
	// TODO: reset session
}
