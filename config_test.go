package hc

import (
	"reflect"
	"testing"

	"github.com/brutella/hc/util"
)

var config = &Config{
	name:         "My MDNS Service",
	id:           "1234",
	servePort:    5010,
	version:      1,
	categoryId:   1,
	state:        1,
	protocol:     "1.0",
	discoverable: true,
	mfiCompliant: false,
}

func TestTxtRecords(t *testing.T) {
	expect := map[string]string{
		"pv": "1.0",
		"id": "1234",
		"c#": "1",
		"s#": "1",
		"sf": "0",
		"ff": "0",
		"md": "My MDNS Service",
		"ci": "1",
	}

	config.discoverable = false

	if x := config.txtRecords(); reflect.DeepEqual(x, expect) == false {
		t.Fatalf("%v != %v", x, expect)
	}
}

func TestVersionUpdate(t *testing.T) {
	storage, err := util.NewTempFileStorage()
	if err != nil {
		t.Fatal(err)
	}

	storage.Set("configHash", []byte("AB"))
	storage.Set("version", []byte("1"))

	config.load(storage)
	config.updateConfigHash([]byte("ABC"))
	config.save(storage)

	if x := config.version; x != 2 {
		t.Fatal(x)
	}

	if x, _ := storage.Get("configHash"); reflect.DeepEqual(x, config.configHash) == false {
		t.Fatal(string(x))
	}

	if x, _ := storage.Get("version"); reflect.DeepEqual(x, []byte("2")) == false {
		t.Fatal(string(x))
	}
}

func TestSetupID(t *testing.T) {
	expected := "1ARVnw=="
	hash := config.setupHash()

	if hash != expected {
		t.Fatalf("generated setup hash is not identical (actual: %s, expected: %s)", hash, expected)
	}

	config.SetupId = "UPDT"

	expected = "PrXxQA=="
	hash = config.setupHash()

	if hash != expected {
		t.Fatalf("generated setup hash is not identical (actual: %s, expected: %s)", hash, expected)
	}

}
