package controller

import(
    "github.com/brutella/hap/model/container"
    "encoding/json"
    "bytes"
    "io"
)

type ContainerController struct {
    container *container.Container
}

func NewContainerController(m *container.Container) *ContainerController {
    return &ContainerController{container: m}
}

func (c *ContainerController) HandleGetAccessories(r io.Reader) (io.Reader, error) {
    result, err := json.Marshal(c.container)
    var b bytes.Buffer
    b.Write(result)
    
    return &b, err
}