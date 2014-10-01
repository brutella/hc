package server

import(
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/common"
    
    "fmt"
    "strings"
    "strconv"
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

func (c *CharacteristicController) GetCharacteristic(accessoryId int, characteristicId int) *model.Characteristic {
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

// string must be in format <accessory id>.<characteristic id>
func ParseAccessoryAndCharacterId(str string) (int, int, error) {
    ids := strings.Split(str, ".")
    if len(ids) != 2 {
        return 0, 0, common.NewErrorf("Could not parse uid %s", str)
    }
    
    aid, err := strconv.Atoi(ids[0])
    cid, err := strconv.Atoi(ids[1])
    
    return aid, cid, err
}