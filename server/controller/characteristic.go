package controller

import(
    "github.com/brutella/hap/model/characteristic"
    "github.com/brutella/hap/model"
    
    "fmt"
)

type Characteristic struct {
    AccessoryId int  `json:"aid"`
    Id int           `json:"iid"`
    Value interface{} `json:"value"`
}

type Characteristics struct {
    Characteristics []Characteristic `json:"characteristics"`
}

func NewCharacteristics() *Characteristics {
    return &Characteristics{
        Characteristics: make([]Characteristic, 0),
    }
}

func (r *Characteristics) AddCharacteristic(c Characteristic) {
    r.Characteristics = append(r.Characteristics, c)
}

type CharacteristicController struct {
    model *model.Model
}

func NewCharacteristicController(m *model.Model) *CharacteristicController {
    return &CharacteristicController{model: m}
}

func (controller *CharacteristicController) HandleGetCharacteristics(aid, cid int) *Characteristics {    
    modelChar := controller.GetCharacteristic(aid, cid)
    if modelChar == nil {
        fmt.Printf("[WARNING] No characteristic found with aid %d and iid %d\n", aid, cid)
    }
    
    chars := NewCharacteristics()
    char := Characteristic{AccessoryId: aid, Id: cid, Value: modelChar.Value}
    chars.AddCharacteristic(char)
    
    return chars
}

func (controller *CharacteristicController) HandleUpdateCharacteristics(chars Characteristics) error {
    for _, c := range chars.Characteristics {
        modelChar := controller.GetCharacteristic(c.AccessoryId, c.Id)
        if modelChar == nil {
            fmt.Printf("[WARNING] Could not find characteristic with aid %d and iid %d\n", c.AccessoryId, c.Id)
            continue
        }
        modelChar.SetValueFromRemote(c.Value)
    }
    
    return nil
}

func (c *CharacteristicController) GetCharacteristic(accessoryId int, characteristicId int) *characteristic.Characteristic {
    for _, a := range c.model.Accessories {
        if a.Id == accessoryId {
            for _, s := range a.Services {
                for _, c :=  range s.Characteristics {
                    if c.Id == characteristicId {
                        return c
                    }
                }
            }
        }
    }
    return nil
}