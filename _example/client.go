package main

import (
	"fmt"
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/netio/pair"
	"io"
	"net/http"
	"os"
)

func sendTLV8(b io.Reader) (io.Reader, error) {
	resp, err := http.Post("http://127.0.0.1:55036/pair-setup", netio.HTTPContentTypePairingTLV8, b)
	return resp.Body, err
}

func main() {
	storage, err := common.NewFileStorage(os.TempDir())
	info := netio.NewBridgeInfo("Test Bridge", "719-47-107", "Matthias H.", storage)
	info.Id = "42:cd:02:57:0d:40"
	db := db.NewDatabaseWithStorage(storage)
	bridge, err := netio.NewBridge(info, db)

	client := pair.NewSetupClientController(bridge, "HomeKit Client")
	pairStartRequest := client.InitialPairingRequest()

	pairStartResponse, err := sendTLV8(pairStartRequest)
	if err != nil {
		fmt.Println(err)
	}

	// 2) S -> C
	pairVerifyRequest, err := pair.HandleReaderForHandler(pairStartResponse, client)
	if err != nil {
		fmt.Println(err)
	}

	// 3) C -> S
	pairVerifyResponse, err := sendTLV8(pairVerifyRequest)
	if err != nil {
		fmt.Println(err)
	}

	// 4) S -> C
	pairKeyRequest, err := pair.HandleReaderForHandler(pairVerifyResponse, client)
	if err != nil {
		fmt.Println(err)
	}

	// 5) C -> S
	pairKeyRespond, err := sendTLV8(pairKeyRequest)
	if err != nil {
		fmt.Println(err)
	}

	// 6) S -> C
	request, err := pair.HandleReaderForHandler(pairKeyRespond, client)
	if err != nil {
		fmt.Println(err)
	}

	if request != nil {
		fmt.Println(request)
	}

	fmt.Println("*** Pairing done ***")

	name := "UnitTest"
	verify := pair.NewVerifyClientController(bridge, name)

	verifyStartRequest := verify.InitialKeyVerifyRequest()
	// 1) C -> S
	verifyStartResponse, err := sendTLV8(verifyStartRequest)
	if err != nil {
		fmt.Println(err)
	}

	// 2) S -> C
	verifyFinishRequest, err := pair.HandleReaderForHandler(verifyStartResponse, verify)
	if err != nil {
		fmt.Println(err)
	}

	// 3) C -> S
	verifyFinishResponse, err := sendTLV8(verifyFinishRequest)
	if err != nil {
		fmt.Println(err)
	}

	// 4) S -> C
	last_request, err := pair.HandleReaderForHandler(verifyFinishResponse, verify)
	if err != nil {
		fmt.Println(err)
	}

	if last_request != nil {
		fmt.Println(last_request)
	}

	fmt.Println("*** Key Verification done ***")
}
