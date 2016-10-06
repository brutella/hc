package controller

import (
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/hap/data"
	"github.com/brutella/hc/log"
	"github.com/gosexy/to"

	"bytes"
	"encoding/json"

	"io"
	"io/ioutil"
	"net"
	"net/url"
	"strings"
)

// CharacteristicController implements the CharacteristicsHandler interface and provides
// read (GET) and write (POST) interfaces to the managed characteristics.
type CharacteristicController struct {
	container *accessory.Container
}

// NewCharacteristicController returns a new characteristic controller.
func NewCharacteristicController(m *accessory.Container) *CharacteristicController {
	return &CharacteristicController{container: m}
}

// HandleGetCharacteristics handles a get characteristic request like `/characteristics?id=1.4,1.5`
func (ctr *CharacteristicController) HandleGetCharacteristics(form url.Values) (io.Reader, error) {
	var b bytes.Buffer
	var chs []data.Characteristic

	// id=1.4,1.5
	paths := strings.Split(form.Get("id"), ",")
	for _, p := range paths {
		if ids := strings.Split(p, "."); len(ids) == 2 {
			aid := to.Int64(ids[0]) // accessory id
			iid := to.Int64(ids[1]) // instance id (= characteristic id)
			c := data.Characteristic{AccessoryID: aid, CharacteristicID: iid}
			if ch := ctr.GetCharacteristic(aid, iid); ch != nil {
				c.Value = ch.Value
			} else {
				c.Status = hap.StatusServiceCommunicationFailure
			}
			chs = append(chs, c)
		}
	}

	result, err := json.Marshal(&data.Characteristics{chs})
	if err != nil {
		log.Info.Panic(err)
	}

	b.Write(result)
	return &b, err
}

// HandleUpdateCharacteristics handles an update characteristic request. The bytes must represent
// a data.Characteristics json.
func (ctr *CharacteristicController) HandleUpdateCharacteristics(r io.Reader, conn net.Conn) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var chars data.Characteristics
	err = json.Unmarshal(b, &chars)
	if err != nil {
		return err
	}

	log.Debug.Println(string(b))

	for _, c := range chars.Characteristics {
		characteristic := ctr.GetCharacteristic(c.AccessoryID, c.CharacteristicID)
		if characteristic == nil {
			log.Info.Printf("Could not find characteristic with aid %d and iid %d\n", c.AccessoryID, c.CharacteristicID)
			continue
		}

		if c.Value != nil {
			characteristic.UpdateValueFromConnection(c.Value, conn)
		}

		if events, ok := c.Events.(bool); ok == true {
			characteristic.SetEventsEnabled(events)
		}
	}

	return err
}

// GetCharacteristic returns the characteristic identified by the accessory id aid and characteristic id iid
func (ctr *CharacteristicController) GetCharacteristic(aid int64, iid int64) *characteristic.Characteristic {
	for _, a := range ctr.container.Accessories {
		if a.GetID() == aid {
			for _, s := range a.GetServices() {
				for _, c := range s.GetCharacteristics() {
					if c.GetID() == iid {
						return c
					}
				}
			}
		}
	}
	return nil
}
