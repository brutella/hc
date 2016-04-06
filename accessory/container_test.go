package accessory

import (
	"testing"
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
	if x := acc1.GetID(); x == 2 {
		t.Fatal(x)
	}
	if x := acc2.GetID(); x == 3 {
		t.Fatal(x)
	}
	if acc1.GetID() == acc2.GetID() {
		t.Fatal("equal ids not allowed")
	}

	c.RemoveAccessory(acc2)

	if is, want := len(c.Accessories), 1; is != want {
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
