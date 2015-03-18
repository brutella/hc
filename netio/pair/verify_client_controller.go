package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"

	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

type VerifyClientController struct {
	client   *netio.Client
	database db.Database
	session  *VerifySession
}

func NewVerifyClientController(client *netio.Client, database db.Database) *VerifyClientController {
	controller := VerifyClientController{
		client:   client,
		database: database,
		session:  NewVerifySession(),
	}

	return &controller
}

func (verify *VerifyClientController) Handle(in common.Container) (common.Container, error) {
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
	case VerifyStepStartResponse:
		out, err = verify.handlePairStepVerifyResponse(in)
	case VerifyStepFinishResponse:
		out, err = verify.handlePairVerifyStepFinishResponse(in)
	default:
		return nil, ErrInvalidVerifyStep(seq)
	}

	return out, err
}

// Client -> Server
// - Public key `A`
func (verify *VerifyClientController) InitialKeyVerifyRequest() io.Reader {
	out := common.NewTLV8Container()
	out.SetByte(TagPairingMethod, 0)
	out.SetByte(TagSequence, VerifyStepStartRequest.Byte())
	out.SetBytes(TagPublicKey, verify.session.PublicKey[:])

	fmt.Println("<-     A:", hex.EncodeToString(out.GetBytes(TagPublicKey)))

	return out.BytesBuffer()
}

// Server -> Client
// - B: server public key
// - encrypted message
//      - username
//      - signature: from server session public key, server name, client session public key
//
// Client -> Server
// - encrypted message
//      - username
//      - signature: from client session public key, server name, server session public key,
func (verify *VerifyClientController) handlePairStepVerifyResponse(in common.Container) (common.Container, error) {
	serverPublicKey := in.GetBytes(TagPublicKey)
	if len(serverPublicKey) != 32 {
		return nil, common.NewErrorf("Invalid server public key size %d", len(serverPublicKey))
	}

	var otherPublicKey [32]byte
	copy(otherPublicKey[:], serverPublicKey)
	verify.session.GenerateSharedKeyWithOtherPublicKey(otherPublicKey)
	verify.session.SetupEncryptionKey([]byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))

	fmt.Println("Client")
	fmt.Println("->   B:", hex.EncodeToString(serverPublicKey))
	fmt.Println("     S:", hex.EncodeToString(verify.session.PrivateKey[:]))
	fmt.Println("Shared:", hex.EncodeToString(verify.session.SharedKey[:]))
	fmt.Println("     K:", hex.EncodeToString(verify.session.EncryptionKey[:]))

	// Decrypt
	data := in.GetBytes(TagEncryptedData)
	message := data[:(len(data) - 16)]
	var mac [16]byte
	copy(mac[:], data[len(message):]) // 16 byte (MAC)

	decryptedBytes, err := crypto.Chacha20DecryptAndPoly1305Verify(verify.session.EncryptionKey[:], []byte("PV-Msg02"), message, mac, nil)
	if err != nil {
		return nil, err
	}

	decryptedIn, err := common.NewTLV8ContainerFromReader(bytes.NewBuffer(decryptedBytes))
	if err != nil {
		return nil, err
	}

	username := decryptedIn.GetString(TagUsername)
	signature := decryptedIn.GetBytes(TagSignature)

	fmt.Println("    Username:", username)
	fmt.Println("   Signature:", hex.EncodeToString(signature))

	// Validate signature
	material := make([]byte, 0)
	material = append(material, verify.session.OtherPublicKey[:]...)
	material = append(material, username...)
	material = append(material, verify.session.PublicKey[:]...)

	entity := verify.database.EntityWithName(username)
	if entity == nil {
		return nil, common.NewErrorf("Server %s is unknown", username)
	}

	if len(entity.PublicKey()) == 0 {
		return nil, common.NewErrorf("No LTPK available for client %s", username)
	}

	if crypto.ValidateED25519Signature(entity.PublicKey(), material, signature) == false {
		return nil, common.NewErrorf("Could not validate signature")
	}

	out := common.NewTLV8Container()
	out.SetByte(TagPairingMethod, PairingMethodDefault)
	out.SetByte(TagSequence, VerifyStepFinishRequest.Byte())

	encryptedOut := common.NewTLV8Container()
	encryptedOut.SetString(TagUsername, verify.client.Id())

	material = make([]byte, 0)
	material = append(material, verify.session.PublicKey[:]...)
	material = append(material, verify.client.Id()...)
	material = append(material, verify.session.OtherPublicKey[:]...)

	signature, err = crypto.ED25519Signature(verify.client.PrivateKey(), material)
	if err != nil {
		return nil, err
	}

	encryptedOut.SetBytes(TagSignature, signature)

	encryptedBytes, mac, _ := crypto.Chacha20EncryptAndPoly1305Seal(verify.session.EncryptionKey[:], []byte("PV-Msg03"), encryptedOut.BytesBuffer().Bytes(), nil)

	out.SetBytes(TagEncryptedData, append(encryptedBytes, mac[:]...))

	return out, nil
}

// Server -> Client
// - only error ocde (optional)
func (verify *VerifyClientController) handlePairVerifyStepFinishResponse(in common.Container) (common.Container, error) {
	code := ErrCode(in.GetByte(TagErrCode))
	if code != ErrCodeNo {
		fmt.Printf("Unexpected error %v\n", code)
	}

	return nil, nil
}
