package controller

import(
    "github.com/brutella/hap/model"
    "encoding/json"
    "bytes"
    "io"
)

type ModelController struct {
    model *model.Model
}

func NewModelController(m *model.Model) *ModelController {
    return &ModelController{model: m}
}

func (c *ModelController) HandleGetAccessories(r io.Reader) (io.Reader, error) {
    result, err := json.Marshal(c.model)
    var b bytes.Buffer
    b.Write(result)
    
    return &b, err
}