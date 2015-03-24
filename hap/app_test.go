package hap

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	// "github.com/brutella/hc/db"
	"github.com/brutella/hc/model"
	// "github.com/brutella/hc/model/accessory"
)

var info = model.Info{
	Name:         "My Switch",
	SerialNumber: "001",
	Manufacturer: "Google",
	Model:        "Switchy",
}

// TODO(brutella) Decide when to update DNS
// func TestDNSUpdate(t *testing.T) {
//     conf := NewConfig()
//     conf.DatabaseDir = os.TempDir()
//
//     app, err := NewApp(conf)
//     assert.Nil(t, err)
//     assert.NotNil(t, app)
//
//     name := app.bridge.Name()
//     dns := db.NewDNS(name, 1, 1)
//     app.Database.SaveDNS(dns)
//     configuration := dns.Configuration()
//
//     sw1 := accessory.NewSwitch(info)
//     app.AddAccessory(sw1.Accessory)
//
//     // DNS must not change because service is not published yet
//     assert.Equal(t, app.Database.DNSWithName(name).Configuration(), configuration)
//
//     sw2 := accessory.NewSwitch(info)
//     app.PerformBatchUpdates(func() {
//         app.AddAccessory(sw2.Accessory)
//         assert.Equal(t, app.Database.DNSWithName(name).Configuration(), configuration)
//     })
//
//     assert.Equal(t, app.Database.DNSWithName(name).Configuration(), configuration + 1)
// }

func TestReachabililty(t *testing.T) {
	conf := NewConfig()
	conf.DatabaseDir = os.TempDir()

	app, err := NewApp(conf)
	assert.Nil(t, err)
	assert.NotNil(t, app)
	assert.False(t, app.IsReachable())
}
