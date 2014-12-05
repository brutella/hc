package app

import (
    "os"
    "testing"
    "github.com/stretchr/testify/assert"
    
    "github.com/brutella/hap/db"
    "github.com/brutella/hap/model"
    "github.com/brutella/hap/model/accessory"
)

var info = model.Info{
    Name: "My Switch",
    SerialNumber: "001",
    Manufacturer: "Google",
    Model: "Switchy",
}

func TestBatchUpdates(t *testing.T) {
    conf := NewConfig()
    conf.DatabaseDir = os.TempDir()
    
    app, err := NewApp(conf)
    assert.Nil(t, err)
    assert.NotNil(t, app)
    
    name := app.bridge.Name()
    dns := db.NewDns(name, 1, 1)
    app.Database.SaveDns(dns)
    configuration := dns.Configuration()
    
    sw1 := accessory.NewSwitch(info)
    sw2 := accessory.NewSwitch(info)
    app.PerformBatchUpdates(func() {
        app.AddAccessory(sw1.Accessory)
        assert.Equal(t, app.Database.DnsWithName(name).Configuration(), configuration)
        app.AddAccessory(sw2.Accessory)
        assert.Equal(t, app.Database.DnsWithName(name).Configuration(), configuration)
    })
    
    // configuration + 1
    assert.Equal(t, app.Database.DnsWithName(name).Configuration(), configuration + 1)
}
