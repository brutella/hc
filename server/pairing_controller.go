package server

import(
    "github.com/brutella/hap"
    "encoding/hex"
    "encoding/json"
    "io"
    "io/ioutil"
    "fmt"
)

/*
{
  guestName: <string>,
  guestPublicKey: <string>
}
*/
type Pairing struct {
    GuestName string `json:"guestName"`
    GuestPublicKey string `json:"guestPublicKey"`
}

type PairingController struct {
    context *hap.Context
}

func NewPairingController(context *hap.Context) *PairingController {
    c := PairingController{
        context: context,
    }
    
    return &c
}

func (c *PairingController) HandleDeletePairings(r io.Reader) (io.Reader, error) {
    client, err := ClientFromRequest(r)
    if client != nil {
        c.context.DeleteClient(client)
    }
    
    return nil, err
}

func (c *PairingController) HandlePostPairings(r io.Reader) (io.Reader, error) {
    client, err := ClientFromRequest(r)
    if client != nil {
        c.context.SaveClient(client)
    }
    
    return nil, err
}

func ClientFromRequest(r io.Reader) (*hap.Client, error) {
    b, _ := ioutil.ReadAll(r)
    var pairing Pairing
    err := json.Unmarshal(b, &pairing)
    
    if err != nil {
        fmt.Println("Could not unmarshal to json", err)
        return nil, err
    }
    
    LTPK, err := hex.DecodeString(pairing.GuestPublicKey)
    
    if err != nil {
        fmt.Println("Could not read guest LTPK", err)
        return nil, err
    }
    
    client := hap.NewClient(pairing.GuestName, LTPK)
    return client, nil
}