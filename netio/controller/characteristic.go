package controller

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/container"
	"github.com/brutella/hc/netio/data"
	"github.com/brutella/log"

	"bytes"
	"encoding/json"

	"io"
	"io/ioutil"
	"net/url"
)

// CharacteristicController implements the CharacteristicsHandler interface.
type CharacteristicController struct {
	container *container.Container
}

func NewCharacteristicController(m *container.Container) *CharacteristicController {
	return &CharacteristicController{container: m}
}

func (controller *CharacteristicController) HandleGetCharacteristics(form url.Values) (io.Reader, error) {
	var b bytes.Buffer
	aid, cid, err := ParseAccessoryAndCharacterId(form.Get("id"))
	containerChar := controller.GetCharacteristic(aid, cid)
	if containerChar == nil {
		log.Printf("[WARN] No characteristic found with aid %d and iid %d\n", aid, cid)
		return &b, nil
	}

	chars := data.NewCharacteristics()
	char := data.Characteristic{AccessoryId: aid, Id: cid, Value: containerChar.GetValue(), Events: containerChar.EventsEnabled()}
	chars.AddCharacteristic(char)

	result, err := json.Marshal(chars)
	if err != nil {
		log.Println("[ERRO]", err)
	}

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

	log.Println("[VERB]", string(b))

	for _, c := range chars.Characteristics {
		containerChar := controller.GetCharacteristic(c.AccessoryId, c.Id)
		if containerChar == nil {
			log.Printf("[ERRO] Could not find characteristic with aid %d and iid %d\n", c.AccessoryId, c.Id)
			continue
		}

		if c.Value != nil {
			containerChar.SetValueFromRemote(c.Value)
		}

		if events, ok := c.Events.(bool); ok == true {
			containerChar.SetEventsEnabled(events)
		}
	}

	return err
}

func (c *CharacteristicController) GetCharacteristic(accessoryId int64, characteristicId int64) model.Characteristic {
	for _, a := range c.container.Accessories {
		if a.GetId() == accessoryId {
			for _, s := range a.GetServices() {
				for _, c := range s.GetCharacteristics() {
					if c.GetId() == characteristicId {
						return c
					}
				}
			}
		}
	}
	return nil
}
