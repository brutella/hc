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
	"fmt"
	"io"
)

// SetupClientController handles pairing with an accessory using SRP.
type SetupClientController struct {
	client   netio.Device
	session  *SetupClientSession
	database db.Database
}

// NewSetupClientController returns a new setup client controller.
func NewSetupClientController(password string, client netio.Device, database db.Database) *SetupClientController {
	session := NewSetupClientSession("Pair-Setup", password)
	controller := SetupClientController{
		client:   client,
		session:  session,
		database: database,
	}
	return &controller
}

// InitialPairingRequest returns the first request the client sends to an accessory to start the paring process.
// The request contains the sequence set to PairStepStartRequest.
func (setup *SetupClientController) InitialPairingRequest() io.Reader {
	out := util.NewTLV8Container()
	out.SetByte(TagPairingMethod, 0)
	out.SetByte(TagSequence, PairStepStartRequest.Byte())

	return out.BytesBuffer()
}

// Handle processes a container to pair (exchange keys) with an accessory.
func (setup *SetupClientController) Handle(in util.Container) (util.Container, error) {
	method := pairMethodType(in.GetByte(TagPairingMethod))

	// It is valid that method is not sent
	// If method is sent then it must be 0x00
	if method != PairingMethodDefault {
		return nil, errInvalidPairMethod(method)
	}

	code := errCode(in.GetByte(TagErrCode))
	if code != ErrCodeNo {
		log.Println("[ERRO]", code)
		return nil, code.Error()
	}

	seq := pairStepType(in.GetByte(TagSequence))

	var out util.Container
	var err error

	switch seq {
	case PairStepStartResponse:
		out, err = setup.handlePairStepStartResponse(in)
	case PairStepVerifyResponse:
		out, err = setup.handlePairStepVerifyResponse(in)
	case PairStepKeyExchangeResponse:
		out, err = setup.handleKeyExchange(in)
	default:
		return nil, errInvalidPairStep(seq)
	}

	return out, err
}

// Server -> Client
// - B: server public key
// - s: salt
//
// Client -> Server
// - A: client public key
// - M1: proof
func (setup *SetupClientController) handlePairStepStartResponse(in util.Container) (util.Container, error) {
	salt := in.GetBytes(TagSalt)
	serverPublicKey := in.GetBytes(TagPublicKey)

	if len(salt) != 16 {
		return nil, fmt.Errorf("Salt is invalid (%d bytes)", len(salt))
	}

	if len(serverPublicKey) != 384 {
		return nil, fmt.Errorf("B is invalid (%d bytes)", len(serverPublicKey))
	}

	fmt.Println("->     B:", hex.EncodeToString(serverPublicKey))
	fmt.Println("->     s:", hex.EncodeToString(salt))

	// Client
	// 1) Receive salt `s` and public key `B` and generates `S` and `A`
	err := setup.session.GenerateKeys(salt, serverPublicKey)
	if err != nil {
		return nil, err
	}
	fmt.Println("        S:", hex.EncodeToString(setup.session.PrivateKey))

	// 2) Send public key `A` and proof `M1`
	publicKey := setup.session.PublicKey // SRP public key
	proof := setup.session.Proof         // M1

	fmt.Println("<-     A:", hex.EncodeToString(publicKey))
	fmt.Println("<-     M1:", hex.EncodeToString(proof))

	out := util.NewTLV8Container()
	out.SetByte(TagPairingMethod, 0)
	out.SetByte(TagSequence, PairStepVerifyRequest.Byte())
	out.SetBytes(TagPublicKey, publicKey)
	out.SetBytes(TagProof, proof)

	return out, nil
}

// Client -> Server
// - A: client public key
// - M1: proof
//
// Server -> client
// - M2: proof
// or
// - auth error
func (setup *SetupClientController) handlePairStepVerifyResponse(in util.Container) (util.Container, error) {
	serverProof := in.GetBytes(TagProof)
	fmt.Println("->     M2:", hex.EncodeToString(serverProof))

	if setup.session.IsServerProofValid(serverProof) == false {
		return nil, fmt.Errorf("M2 %s is invalid", hex.EncodeToString(serverProof))
	}

	err := setup.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
	if err != nil {
		return nil, err
	}

	fmt.Println("        K:", hex.EncodeToString(setup.session.EncryptionKey[:]))

	// 2) Send username, LTPK, signature as encrypted message
	hash, err := hkdf.Sha512(setup.session.PrivateKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
	var material []byte
	material = append(material, hash[:]...)
	material = append(material, setup.client.Name()...)
	material = append(material, setup.client.PublicKey()...)

	signature, err := crypto.ED25519Signature(setup.client.PrivateKey(), material)
	if err != nil {
		return nil, err
	}

	encryptedOut := util.NewTLV8Container()
	encryptedOut.SetString(TagUsername, setup.client.Name())
	encryptedOut.SetBytes(TagPublicKey, []byte(setup.client.PublicKey()))
	encryptedOut.SetBytes(TagSignature, []byte(signature))

	encryptedBytes, tag, err := chacha20poly1305.EncryptAndSeal(setup.session.EncryptionKey[:], []byte("PS-Msg05"), encryptedOut.BytesBuffer().Bytes(), nil)
	if err != nil {
		return nil, err
	}

	out := util.NewTLV8Container()
	out.SetByte(TagPairingMethod, 0)
	out.SetByte(TagSequence, PairStepKeyExchangeRequest.Byte())
	out.SetBytes(TagEncryptedData, append(encryptedBytes, tag[:]...))

	fmt.Println("<-   Encrypted:", hex.EncodeToString(out.GetBytes(TagEncryptedData)))

	return out, nil
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
// - encrpyted tlv8: bridge LTPK, bridge name, signature (of hash `H2`, bridge name, LTPK)
func (setup *SetupClientController) handleKeyExchange(in util.Container) (util.Container, error) {
	data := in.GetBytes(TagEncryptedData)
	message := data[:(len(data) - 16)]
	var mac [16]byte
	copy(mac[:], data[len(message):]) // 16 byte (MAC)
	fmt.Println("->     Message:", hex.EncodeToString(message))
	fmt.Println("->     MAC:", hex.EncodeToString(mac[:]))

	decrypted, err := chacha20poly1305.DecryptAndVerify(setup.session.EncryptionKey[:], []byte("PS-Msg06"), message, mac, nil)

	if err != nil {
		fmt.Println(err)
	} else {
		decryptedBuf := bytes.NewBuffer(decrypted)
		in, err := util.NewTLV8ContainerFromReader(decryptedBuf)
		if err != nil {
			fmt.Println(err)
		}

		username := in.GetString(TagUsername)
		ltpk := in.GetBytes(TagPublicKey)
		signature := in.GetBytes(TagSignature)
		fmt.Println("->     Username:", username)
		fmt.Println("->     LTPK:", hex.EncodeToString(ltpk))
		fmt.Println("->     Signature:", hex.EncodeToString(signature))

		entity := db.NewEntity(username, ltpk, nil)
		err = setup.database.SaveEntity(entity)
		if err != nil {
			fmt.Println("[ERRO]", err)
		}
	}

	return nil, err
}
