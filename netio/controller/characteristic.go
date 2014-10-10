package controller

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/netio/data"
    
    "fmt"
    "encoding/json"
    "bytes"

    "net/url"
    "io"
    "io/ioutil"
)

type CharacteristicController struct {
    model *model.Model
}

func NewCharacteristicController(m *model.Model) *CharacteristicController {
    return &CharacteristicController{model: m}
}

func (controller *CharacteristicController) HandleGetCharacteristics(form url.Values) (io.Reader, error) {    
    aid, cid, err := ParseAccessoryAndCharacterId(form.Get("id"))
    modelChar := controller.GetCharacteristic(aid, cid)
    if modelChar == nil {
        fmt.Printf("[WARNING] No characteristic found with aid %d and iid %d\n", aid, cid)
    }
    
    chars := data.NewCharacteristics()
    char := data.Characteristic{AccessoryId: aid, Id: cid, Value: modelChar.GetValue()}
    chars.AddCharacteristic(char)
    
    result, err := json.Marshal(chars)
    if err != nil {
        fmt.Println(err)
    }
    
    var b bytes.Buffer
    b.Write(result)
    return &b, err
}

func (controller *CharacteristicController) HandleUpdateCharacteristics(r io.Reader) error {
    b, err := ioutil.ReadAll(r)
    if err != nil {
        return err
    }
    
    var chars data.Characteristics
    err = json.Unmarshal(b, &chars)
    if err != nil {
        return err
    }
    
    for _, c := range chars.Characteristics {
        modelChar := controller.GetCharacteristic(c.AccessoryId, c.Id)
        if modelChar == nil {
            fmt.Printf("[WARNING] Could not find characteristic with aid %d and iid %d\n", c.AccessoryId, c.Id)
            continue
        }
        modelChar.SetValueFromRemote(c.Value)
    }
    
    return err
}

func (c *CharacteristicController) GetCharacteristic(accessoryId int, characteristicId int) model.Characteristic {
    for _, a := range c.model.Accessories {
        if a.Id() == accessoryId {
            for _, s := range a.GetServices() {
                for _, c :=  range s.GetCharacteristics() {
                    if c.Id() == characteristicId {
                        return c
                    }
                }
            }
        }
    }
    return nil
}