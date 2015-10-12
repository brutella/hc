package pair

import (
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/crypto/chacha20poly1305"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/util"

	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

// VerifyClientController verifies the stored accessory public key and negotiates a shared secret
// which is used encrypt the upcoming session.
//
// Verification fails when the accessory is not known, the public key for the accessory was not found,
// or the packet's seal could not be verified.
type VerifyClientController struct {
	client   netio.Device
	database db.Database
	session  *VerifySession
}

// NewVerifyClientController returns a new verify client controller.
func NewVerifyClientController(client netio.Device, database db.Database) *VerifyClientController {
	controller := VerifyClientController{
		client:   client,
		database: database,
		session:  NewVerifySession(),
	}

	return &controller
}

// Handle processes a container to verify if an accessory is paired correctly.
func (verify *VerifyClientController) Handle(in util.Container) (util.Container, error) {
	var out util.Container
	var err error

	method := PairMethodType(in.GetByte(TagPairingMethod))

	// It is valid that method is not sent
	// If method is sent then it must be 0x00
	if method != PairingMethodDefault {
		return nil, errInvalidPairMethod(method)
	}

	seq := VerifyStepType(in.GetByte(TagSequence))
	switch seq {
	case VerifyStepStartResponse:
		out, err = verify.handlePairStepVerifyResponse(in)
	case VerifyStepFinishResponse:
		out, err = verify.handlePairVerifyStepFinishResponse(in)
	default:
		return nil, errInvalidVerifyStep(seq)
	}

	return out, err
}

// InitialKeyVerifyRequest returns the first request the client sends to an accessory to start the paring verifcation process.
// The request contains the client public key and sequence set to VerifyStepStartRequest.
func (verify *VerifyClientController) InitialKeyVerifyRequest() io.Reader {
	out := util.NewTLV8Container()
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
func (verify *VerifyClientController) handlePairStepVerifyResponse(in util.Container) (util.Container, error) {
	serverPublicKey := in.GetBytes(TagPublicKey)
	if len(serverPublicKey) != 32 {
		return nil, fmt.Errorf("Invalid server public key size %d", len(serverPublicKey))
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

	decryptedBytes, err := chacha20poly1305.DecryptAndVerify(verify.session.EncryptionKey[:], []byte("PV-Msg02"), message, mac, nil)
	if err != nil {
		return nil, err
	}

	decryptedIn, err := util.NewTLV8ContainerFromReader(bytes.NewBuffer(decryptedBytes))
	if err != nil {
		return nil, err
	}

	username := decryptedIn.GetString(TagUsername)
	signature := decryptedIn.GetBytes(TagSignature)

	fmt.Println("    Username:", username)
	fmt.Println("   Signature:", hex.EncodeToString(signature))

	// Validate signature
	var material []byte
	material = append(material, verify.session.OtherPublicKey[:]...)
	material = append(material, username...)
	material = append(material, verify.session.PublicKey[:]...)

	var entity db.Entity
	if entity, err = verify.database.EntityWithName(username); err != nil {
		return nil, fmt.Errorf("Server %s is unknown", username)
	}

	if len(entity.PublicKey) == 0 {
		return nil, fmt.Errorf("No LTPK available for client %s", username)
	}

	if crypto.ValidateED25519Signature(entity.PublicKey, material, signature) == false {
		return nil, fmt.Errorf("Could not validate signature")
	}

	out := util.NewTLV8Container()
	out.SetByte(TagSequence, VerifyStepFinishRequest.Byte())

	encryptedOut := util.NewTLV8Container()
	encryptedOut.SetString(TagUsername, verify.client.Name())

	material = make([]byte, 0)
	material = append(material, verify.session.PublicKey[:]...)
	material = append(material, verify.client.Name()...)
	material = append(material, verify.session.OtherPublicKey[:]...)

	signature, err = crypto.ED25519Signature(verify.client.PrivateKey(), material)
	if err != nil {
		return nil, err
	}

	encryptedOut.SetBytes(TagSignature, signature)

	encryptedBytes, mac, _ := chacha20poly1305.EncryptAndSeal(verify.session.EncryptionKey[:], []byte("PV-Msg03"), encryptedOut.BytesBuffer().Bytes(), nil)

	out.SetBytes(TagEncryptedData, append(encryptedBytes, mac[:]...))

	return out, nil
}

// Server -> Client
// - only error ocde (optional)
func (verify *VerifyClientController) handlePairVerifyStepFinishResponse(in util.Container) (util.Container, error) {
	code := errCode(in.GetByte(TagErrCode))
	if code != ErrCodeNo {
		fmt.Printf("Unexpected error %v\n", code)
	}

	return nil, nil
}
