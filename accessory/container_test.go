package accessory

import (
	"reflect"
	"testing"

	"github.com/brutella/hc/service"
)

var info = Info{
	Name:         "Accessory1",
	SerialNumber: "001",
	Manufacturer: "Google",
	Model:        "Accessory",
}

func TestContainer(t *testing.T) {
	acc1 := New(info, TypeOther)
	info.Name = "Accessory2"
	acc2 := New(info, TypeOther)

	c := NewContainer()
	c.AddAccessory(acc1)
	c.AddAccessory(acc2)

	if is, want := len(c.Accessories), 2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if x := acc1.ID; x == 2 {
		t.Fatal(x)
	}
	if x := acc2.ID; x == 3 {
		t.Fatal(x)
	}
	if acc1.ID == acc2.ID {
		t.Fatal("equal ids not allowed")
	}

	c.RemoveAccessory(acc2)

	if is, want := len(c.Accessories), 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestContainerCustomID(t *testing.T) {
	info1 := info
	info1.ID = 99
	info2 := info
	info2.ID = 15
	info3 := info
	info3.ID = 15

	acc1 := New(info1, TypeOther)
	info2.Name = "Accessory2"
	acc2 := New(info2, TypeOther)
	info3.Name = "Accessory3"
	acc3 := New(info3, TypeOther)

	c := NewContainer()
	c.AddAccessory(acc1)
	c.AddAccessory(acc2)
	c.AddAccessory(acc3)

	if is, want := len(c.Accessories), 3; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := acc1.ID, int64(99); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := acc2.ID, int64(15); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := acc3.ID, int64(100); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if acc2.ID == acc3.ID {
		t.Fatal("equal ids not allowed")
	}

	c.RemoveAccessory(acc2)

	if is, want := len(c.Accessories), 2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestAccessoryCount(t *testing.T) {
	accessory := New(info, TypeOther)
	c := NewContainer()
	c.AddAccessory(accessory)

	if is, want := len(c.Accessories), 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.RemoveAccessory(accessory)

	if is, want := len(c.Accessories), 0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestAccessoryType(t *testing.T) {
	a1 := New(info, TypeLightbulb)
	a2 := New(info, TypeSwitch)

	c := NewContainer()
	c.AddAccessory(a1)

	if is, want := c.AccessoryType(), TypeLightbulb; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.AddAccessory(a2)

	if is, want := c.AccessoryType(), TypeBridge; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestContentHash(t *testing.T) {
	acc := New(info, TypeLightbulb)
	c := NewContainer()
	c.AddAccessory(acc)

	hash := c.ContentHash()

	acc.Info.Name.SetValue("Test Value")

	// Hash ignores the value field and should therefore be the same
	if is, want := c.ContentHash(), hash; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}

	acc.AddService(service.New(service.TypeLightbulb))

	// Hash changes when accessories/services/characteristics are added
	if is, want := c.ContentHash(), hash; reflect.DeepEqual(is, want) == true {
		t.Fatalf("%v should not be %v", is, want)
	}
}
