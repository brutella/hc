package controller

import (
	"bytes"
	"encoding/json"
	"github.com/brutella/hc/model/container"
	"io"
)

// ContainerController implements the AccessoriesHandler interface.
type ContainerController struct {
	container *container.Container
}

// NewContainerController returns a controller for the argument container.
func NewContainerController(m *container.Container) *ContainerController {
	return &ContainerController{container: m}
}

func (c *ContainerController) HandleGetAccessories(r io.Reader) (io.Reader, error) {
	result, err := json.Marshal(c.container)
	return bytes.NewBuffer(result), err
}
