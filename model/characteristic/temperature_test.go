package characteristic

import (
	"reflect"
	"testing"
)

func TestTemperatureCharacteristic(t *testing.T) {
	temp := NewCurrentTemperatureCharacteristic(20.2, 0, 100, 1, "celsius")

	if is, want := temp.Temperature(), 20.2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := temp.MinTemperature(), 0.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := temp.MaxTemperature(), 100.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := temp.MinStepTemperature(), 1.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	temp.SetTemperature(10.1)

	if is, want := temp.Temperature(), 10.1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCurrentTemperatureCharacteristic(t *testing.T) {
	temp := NewCurrentTemperatureCharacteristic(20.2, 0, 100, 1, "celsius")

	if is, want := temp.Type, TypeTemperatureCurrent; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if x := temp.Permissions; reflect.DeepEqual(x, PermsRead()) == false {
		t.Fatal(x)
	}
}

func TestTargetTemperatureCharacteristic(t *testing.T) {
	temp := NewTargetTemperatureCharacteristic(20.2, 0, 100, 1, "celsius")

	if is, want := temp.Type, TypeTemperatureTarget; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if x := temp.Permissions; reflect.DeepEqual(x, PermsAll()) == false {
		t.Fatal(x)
	}
}
