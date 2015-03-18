package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"
	"github.com/brutella/log"

	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

type SetupClientController struct {
	client   *netio.Client
	session  *SetupClientSession
	database db.Database
}

func NewSetupClientController(password string, client *netio.Client, database db.Database) *SetupClientController {
	session := NewSetupClientSession("Pair-Setup", password)
	controller := SetupClientController{
		client:   client,
		session:  session,
		database: database,
	}
	return &controller
}

func (setup *SetupClientController) InitialPairingRequest() io.Reader {
	out := common.NewTLV8Container()
	out.SetByte(TagPairingMethod, 0)
	out.SetByte(TagSequence, PairStepStartRequest.Byte())

	return out.BytesBuffer()
}

func (setup *SetupClientController) Handle(in common.Container) (common.Container, error) {
	method := PairMethodType(in.GetByte(TagPairingMethod))

	// It is valid that method is not sent
	// If method is sent then it must be 0x00
	if method != PairingMethodDefault {
		return nil, ErrInvalidPairMethod(method)
	}

	code := ErrCode(in.GetByte(TagErrCode))
	if code != ErrCodeNo {
		log.Println("[ERRO]", code)
		return nil, code.Error()
	}

	seq := PairStepType(in.GetByte(TagSequence))

	var out common.Container
	var err error

	switch seq {
	case PairStepStartResponse:
		out, err = setup.handlePairStepStartResponse(in)
	case PairStepVerifyResponse:
		out, err = setup.handlePairStepVerifyResponse(in)
	case PairStepKeyExchangeResponse:
		out, err = setup.handleKeyExchange(in)
	default:
		return nil, ErrInvalidPairStep(seq)
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
func (setup *SetupClientController) handlePairStepStartResponse(in common.Container) (common.Container, error) {
	salt := in.GetBytes(TagSalt)
	serverPublicKey := in.GetBytes(TagPublicKey)

	if len(salt) != 16 {
		return nil, common.NewErrorf("Salt is invalid (%d bytes)", len(salt))
	}

	if len(serverPublicKey) != 384 {
		return nil, common.NewErrorf("B is invalid (%d bytes)", len(serverPublicKey))
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

	out := common.NewTLV8Container()
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
func (setup *SetupClientController) handlePairStepVerifyResponse(in common.Container) (common.Container, error) {
	serverProof := in.GetBytes(TagProof)
	fmt.Println("->     M2:", hex.EncodeToString(serverProof))

	if setup.session.IsServerProofValid(serverProof) == false {
		return nil, common.NewErrorf("M2 %s is invalid", hex.EncodeToString(serverProof))
	}

	err := setup.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
	if err != nil {
		return nil, err
	}

	fmt.Println("        K:", hex.EncodeToString(setup.session.EncryptionKey[:]))

	// 2) Send username, LTPK, signature as encrypted message
	hash, err := crypto.HKDF_SHA512(setup.session.PrivateKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
	material := make([]byte, 0)
	material = append(material, hash[:]...)
	material = append(material, setup.client.PairUsername()...)
	material = append(material, setup.client.PairPublicKey()...)

	signature, err := crypto.ED25519Signature(setup.client.PairPrivateKey(), material)
	if err != nil {
		return nil, err
	}

	encryptedOut := common.NewTLV8Container()
	encryptedOut.SetString(TagUsername, setup.client.PairUsername())
	encryptedOut.SetBytes(TagPublicKey, []byte(setup.client.PairPublicKey()))
	encryptedOut.SetBytes(TagSignature, []byte(signature))

	encryptedBytes, tag, err := crypto.Chacha20EncryptAndPoly1305Seal(setup.session.EncryptionKey[:], []byte("PS-Msg05"), encryptedOut.BytesBuffer().Bytes(), nil)
	if err != nil {
		return nil, err
	}

	out := common.NewTLV8Container()
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
func (setup *SetupClientController) handleKeyExchange(in common.Container) (common.Container, error) {
	data := in.GetBytes(TagEncryptedData)
	message := data[:(len(data) - 16)]
	var mac [16]byte
	copy(mac[:], data[len(message):]) // 16 byte (MAC)
	fmt.Println("->     Message:", hex.EncodeToString(message))
	fmt.Println("->     MAC:", hex.EncodeToString(mac[:]))

	decrypted, err := crypto.Chacha20DecryptAndPoly1305Verify(setup.session.EncryptionKey[:], []byte("PS-Msg06"), message, mac, nil)

	if err != nil {
		fmt.Println(err)
	} else {
		decrypted_buffer := bytes.NewBuffer(decrypted)
		in, err := common.NewTLV8ContainerFromReader(decrypted_buffer)
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
