package main

import (
	"fmt"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/netio"
	"github.com/brutella/hc/netio/pair"
	"io"
	"log"
	"net/http"
)

func pairSetup(b io.Reader) (io.Reader, error) {
	return sendTLV8(b, "pair-setup")
}

func pairVerify(b io.Reader) (io.Reader, error) {
	return sendTLV8(b, "pair-verify")
}

func sendTLV8(b io.Reader, endpoint string) (io.Reader, error) {
	url := fmt.Sprintf("http://127.0.0.1:49624/%s", endpoint)
	resp, err := http.Post(url, netio.HTTPContentTypePairingTLV8, b)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Invalid status code %v", resp.StatusCode)
	}
	return resp.Body, err
}

func main() {
	database, _ := db.NewDatabase("./data")
	c, _ := netio.NewClient("Golang Client", database)
	client := pair.NewSetupClientController("740-51-881", c, database)
	pairStartRequest := client.InitialPairingRequest()

	pairStartResponse, err := pairSetup(pairStartRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 2) S -> C
	pairVerifyRequest, err := pair.HandleReaderForHandler(pairStartResponse, client)
	if err != nil {
		log.Fatal(err)
	}

	// 3) C -> S
	pairVerifyResponse, err := pairSetup(pairVerifyRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 4) S -> C
	pairKeyRequest, err := pair.HandleReaderForHandler(pairVerifyResponse, client)
	if err != nil {
		log.Fatal(err)
	}

	// 5) C -> S
	pairKeyRespond, err := pairSetup(pairKeyRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 6) S -> C
	request, err := pair.HandleReaderForHandler(pairKeyRespond, client)
	if err != nil {
		log.Fatal(err)
	}

	if request != nil {
		log.Println(request)
	}

	log.Println("*** Pairing done ***")

	verify := pair.NewVerifyClientController(c, database)

	verifyStartRequest := verify.InitialKeyVerifyRequest()
	// 1) C -> S
	verifyStartResponse, err := pairVerify(verifyStartRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 2) S -> C
	verifyFinishRequest, err := pair.HandleReaderForHandler(verifyStartResponse, verify)
	if err != nil {
		log.Fatal(err)
	}

	// 3) C -> S
	verifyFinishResponse, err := pairVerify(verifyFinishRequest)
	if err != nil {
		log.Fatal(err)
	}

	// 4) S -> C
	last_request, err := pair.HandleReaderForHandler(verifyFinishResponse, verify)
	if err != nil {
		log.Fatal(err)
	}

	if last_request != nil {
		log.Println(last_request)
	}

	log.Println("*** Key Verification done ***")
}
