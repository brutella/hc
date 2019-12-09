package util

import (
	"github.com/brutella/hc/accessory"
	"testing"
)

func TestBadSetupCode(t *testing.T) {
	setupCode := "BAD"
	setupID := "ERIC"
	actual, err := XHMURI(setupCode, setupID, uint8(accessory.TypeLightbulb), []SetupFlag{SetupFlagIP})
	if err == nil {
		t.Errorf("Failed to generate X-HM URI: %s", err)
	}
	if actual != "" {
		t.Errorf("X-HM URI should have been empty, but it was: %v", actual)
	}
}

func TestXHKURL(t *testing.T) {
	setupCode := "102-93-847"
	setupID := "ERIC"
	actual, err := XHMURI(setupCode, setupID, uint8(accessory.TypeLightbulb), []SetupFlag{SetupFlagIP})
	if err != nil {
		t.Errorf("Failed to generate X-HM URI: %s", err)
	}
	expected := "X-HM://00526Q9UFERIC"
	if expected != actual {
		t.Errorf("Expected: '%s', Actual: '%s'", expected, actual)
	}
}

func TestXHMURLWithBluetooth(t *testing.T) {
	setupCode := "102-93-847"
	setupID := "ERIC"
	actual, err := XHMURI(setupCode, setupID, uint8(accessory.TypeLightbulb), []SetupFlag{SetupFlagIP, SetupFlagBTLE})
	if err != nil {
		t.Errorf("Failed to generate X-HM URI: %s", err)
	}
	expected := "X-HM://005B2DA3BERIC"
	if expected != actual {
		t.Errorf("Expected: '%s', Actual: '%s'", expected, actual)
	}
}

func TestGarageDoor(t *testing.T) {
	setupCode := "102-93-847"
	setupID := "ERIC"
	actual, err := XHMURI(setupCode, setupID, uint8(accessory.TypeGarageDoorOpener), []SetupFlag{SetupFlagIP})
	if err != nil {
		t.Errorf("Failed to generate X-HM URI: %s", err)
	}
	expected := "X-HM://0042O68UVERIC"
	if expected != actual {
		t.Errorf("Expected: '%s', Actual: '%s'", expected, actual)
	}
}
