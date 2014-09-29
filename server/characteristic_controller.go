package server

import(
    "github.com/brutella/hap"
    "github.com/brutella/hap/model"
    "encoding/json"
    "net/url"
    "bytes"
    "io"
    "fmt"
    "strings"
    "strconv"
)

type Characteristic struct {
    AccessoryId int  `json:"aid"`
    Id int           `json:"iid"`
    Value interface{} `json:"value"`
}

type CharacteristicsResponse struct {
    Characteristics []Characteristic `json:"characteristics"`
}

func NewCharacteristicsResponse() *CharacteristicsResponse {
    return &CharacteristicsResponse{
        Characteristics: make([]Characteristic, 0),
    }
}

func (r *CharacteristicsResponse) AddCharacteristic(c Characteristic) {
    r.Characteristics = append(r.Characteristics, c)
}

type CharacteristicController struct {
    model *model.Model
}

func NewCharacteristicController(m *model.Model) *CharacteristicController {
    return &CharacteristicController{model: m}
}

func (controller *CharacteristicController) HandleGetCharacteristics(form url.Values) (io.Reader, error) {
    // todo parse
    aid, cid, err := ParseAccessoryAndCharacterId(form.Get("id"))
    if err != nil {
        return nil, err
    }
    
    value := controller.ValueForCharacteristics(aid, cid)
    if value == nil {
        fmt.Printf("[WARNING] No characteristic found with aid %d and iid %d\n", aid, cid)
    }
    
    chars := NewCharacteristicsResponse()
    char := Characteristic{AccessoryId: aid, Id: cid, Value: value}
    chars.AddCharacteristic(char)
    
    result, err := json.Marshal(chars)
    var b bytes.Buffer
    b.Write(result)
    
    return &b, err
}

func (c *CharacteristicController) ValueForCharacteristics(accessoryId int, characteristicId int) interface{} {
    for _, a := range c.model.Accessories {
        if a.Id == accessoryId {
            for _, s := range a.Services {
                for _, c :=  range s.Characteristics {
                    if c.Id == characteristicId {
                        return c.Value
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
        return 0, 0, hap.NewErrorf("Could not parse uid %s", str)
    }
    
    aid, err := strconv.Atoi(ids[0])
    cid, err := strconv.Atoi(ids[1])
    
    return aid, cid, err
}