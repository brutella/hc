package controller

import (
	"bytes"
	"encoding/json"
	"github.com/brutella/hap/model/container"
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
	return bytes.NewBuffer(result), err
}
