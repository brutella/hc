package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/crypto/chacha20poly1305"
	"github.com/brutella/hc/crypto/hkdf"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"
	"github.com/brutella/log"

	"bytes"
	"encoding/hex"
	"errors"
)

// SetupServerController handles pairing with a cliet using SRP.
// The entity has to known the bridge password to successfully pair.
// When pairing was successful, the entity's public key (refered as ltpk - long term public key)
// is stored in the database.
//
// Pairing may fail because the password is wrong or the key exchange failed (e.g. packet seals or SRP key authenticator is wrong, ...).
type SetupServerController struct {
	bridge   *netio.Bridge
	session  *SetupServerSession
	step     pairStepType
	database db.Database
}

// NewSetupServerController returns a new pair setup controller.
func NewSetupServerController(bridge *netio.Bridge, database db.Database) (*SetupServerController, error) {
	if len(bridge.PairPrivateKey()) == 0 {
		return nil, errors.New("no private key for pairing available")
	}

	session, err := NewSetupServerSession(bridge.PairUsername(), bridge.Password())
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

// Handle processes a container to pair (exchange keys) with a client.
func (setup *SetupServerController) Handle(in common.Container) (out common.Container, err error) {
	method := pairMethodType(in.GetByte(TagPairingMethod))

	// It is valid that pair method is not sent
	// If method set then it must be 0x00
	if method != PairingMethodDefault {
		return nil, errInvalidPairMethod(method)
	}

	seq := pairStepType(in.GetByte(TagSequence))

	switch seq {
	case PairStepStartRequest:
		if setup.step != PairStepWaiting {
			setup.reset()
			return nil, errInvalidInternalPairStep(setup.step)
		}

		out, err = setup.handlePairStart(in)
	case PairStepVerifyRequest:
		if setup.step != PairStepStartResponse {
			setup.reset()
			return nil, errInvalidInternalPairStep(setup.step)
		}

		out, err = setup.handlePairVerify(in)
	case PairStepKeyExchangeRequest:
		if setup.step != PairStepVerifyResponse {
			setup.reset()
			return nil, errInvalidInternalPairStep(setup.step)
		}

		out, err = setup.handleKeyExchange(in)
	default:
		return nil, errInvalidPairStep(seq)
	}

	return out, err
}

// Client -> Server
// - Auth start
//
// Server -> Client
// - B: server public key
// - s: salt
func (setup *SetupServerController) handlePairStart(in common.Container) (common.Container, error) {
	out := common.NewTLV8Container()
	setup.step = PairStepStartResponse

	out.SetByte(TagSequence, setup.step.Byte())
	out.SetBytes(TagPublicKey, setup.session.PublicKey)
	out.SetBytes(TagSalt, setup.session.Salt)

	log.Println("[VERB] <-     B:", hex.EncodeToString(out.GetBytes(TagPublicKey)))
	log.Println("[VERB] <-     s:", hex.EncodeToString(out.GetBytes(TagSalt)))

	return out, nil
}

// Client -> Server
// - A: entity public key
// - M1: proof
//
// Server -> entity
// - M2: proof
// or
// - auth error
func (setup *SetupServerController) handlePairVerify(in common.Container) (common.Container, error) {
	setup.step = PairStepVerifyResponse
	out := common.NewTLV8Container()
	out.SetByte(TagSequence, setup.step.Byte())

	clientPublicKey := in.GetBytes(TagPublicKey)
	log.Println("[VERB] ->     A:", hex.EncodeToString(clientPublicKey))

	err := setup.session.SetupPrivateKeyFromClientPublicKey(clientPublicKey)
	if err != nil {
		return nil, err
	}

	clientProof := in.GetBytes(TagProof)
	log.Println("[VERB] ->     M1:", hex.EncodeToString(clientProof))

	proof, err := setup.session.ProofFromClientProof(clientProof)
	if err != nil || len(proof) == 0 { // proof `M1` is wrong
		log.Println("[WARN] Proof M1 is wrong")
		setup.reset()
		out.SetByte(TagErrCode, ErrCodeAuthenticationFailed.Byte()) // return error 2
	} else {
		log.Println("[INFO] Proof M1 is valid")
		err := setup.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
		if err != nil {
			return nil, err
		}

		// Return proof `M2`
		out.SetBytes(TagProof, proof)
	}

	log.Println("[VERB] <-     M2:", hex.EncodeToString(out.GetBytes(TagProof)))
	log.Println("[VERB]         S:", hex.EncodeToString(setup.session.PrivateKey))
	log.Println("[VERB]         K:", hex.EncodeToString(setup.session.EncryptionKey[:]))

	return out, nil
}

// Client -> Server
// - encrypted tlv8: entity ltpk, entity name and signature (of H, entity name, ltpk)
// - auth tag (mac)
//
// Server
// - Validate signature of encrpyted tlv8
// - Read and store entity ltpk and name
//
// Server -> Client
// - encrpyted tlv8: bridge ltpk, bridge name, signature (of hash, bridge name, ltpk)
func (setup *SetupServerController) handleKeyExchange(in common.Container) (common.Container, error) {
	out := common.NewTLV8Container()

	setup.step = PairStepKeyExchangeResponse

	out.SetByte(TagSequence, setup.step.Byte())

	data := in.GetBytes(TagEncryptedData)
	message := data[:(len(data) - 16)]
	var mac [16]byte
	copy(mac[:], data[len(message):]) // 16 byte (MAC)
	log.Println("[VERB] ->     Message:", hex.EncodeToString(message))
	log.Println("[VERB] ->     MAC:", hex.EncodeToString(mac[:]))

	decrypted, err := chacha20poly1305.DecryptAndVerify(setup.session.EncryptionKey[:], []byte("PS-Msg05"), message, mac, nil)

	if err != nil {
		setup.reset()
		log.Println("[ERRO]", err)
		out.SetByte(TagErrCode, ErrCodeUnknown.Byte()) // return error 1
	} else {
		decryptedBuf := bytes.NewBuffer(decrypted)
		in, err := common.NewTLV8ContainerFromReader(decryptedBuf)
		if err != nil {
			return nil, err
		}

		username := in.GetString(TagUsername)
		clientltpk := in.GetBytes(TagPublicKey)
		signature := in.GetBytes(TagSignature)
		log.Println("[VERB] ->     Username:", username)
		log.Println("[VERB] ->     ltpk:", hex.EncodeToString(clientltpk))
		log.Println("[VERB] ->     Signature:", hex.EncodeToString(signature))

		// Calculate hash `H`
		hash, _ := hkdf.Sha512(setup.session.PrivateKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
		var material []byte
		material = append(material, hash[:]...)
		material = append(material, []byte(username)...)
		material = append(material, clientltpk...)

		if crypto.ValidateED25519Signature(clientltpk, material, signature) == false {
			log.Println("[WARN] ed25519 signature is invalid")
			setup.reset()
			out.SetByte(TagErrCode, ErrCodeAuthenticationFailed.Byte()) // return error 2
		} else {
			log.Println("[VERB] ed25519 signature is valid")
			// Store entity ltpk and name
			entity := db.NewEntity(username, clientltpk, nil)
			setup.database.SaveEntity(entity)
			log.Printf("[INFO] Stored ltpk '%s' for entity '%s'\n", hex.EncodeToString(clientltpk), username)

			ltpk := setup.bridge.PairPublicKey()
			ltsk := setup.bridge.PairPrivateKey()

			// Send username, ltpk, signature as encrypted message
			hash, err := hkdf.Sha512(setup.session.PrivateKey, []byte("Pair-Setup-Accessory-Sign-Salt"), []byte("Pair-Setup-Accessory-Sign-Info"))
			material = make([]byte, 0)
			material = append(material, hash[:]...)
			material = append(material, []byte(setup.session.Username)...)
			material = append(material, ltpk...)

			signature, err := crypto.ED25519Signature(ltsk, material)
			if err != nil {
				log.Fatal(err)
				return nil, err
			}

			tlvPairKeyExchange := common.NewTLV8Container()
			tlvPairKeyExchange.SetBytes(TagUsername, setup.session.Username)
			tlvPairKeyExchange.SetBytes(TagPublicKey, ltpk)
			tlvPairKeyExchange.SetBytes(TagSignature, []byte(signature))

			log.Println("[VERB] <-     Username:", tlvPairKeyExchange.GetString(TagUsername))
			log.Println("[VERB] <-     ltpk:", hex.EncodeToString(tlvPairKeyExchange.GetBytes(TagPublicKey)))
			log.Println("[VERB] <-     Signature:", hex.EncodeToString(tlvPairKeyExchange.GetBytes(TagSignature)))

			encrypted, mac, _ := chacha20poly1305.EncryptAndSeal(setup.session.EncryptionKey[:], []byte("PS-Msg06"), tlvPairKeyExchange.BytesBuffer().Bytes(), nil)
			out.SetByte(TagPairingMethod, 0)
			out.SetByte(TagSequence, PairStepKeyExchangeRequest.Byte())
			out.SetBytes(TagEncryptedData, append(encrypted, mac[:]...))

			setup.reset()
		}
	}

	return out, nil
}

func (setup *SetupServerController) reset() {
	setup.step = PairStepWaiting
	// TODO: reset session
}
