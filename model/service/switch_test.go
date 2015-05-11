package service

import (
	"github.com/brutella/hc/model/characteristic"
	"reflect"
	"testing"
)

func TestSwitch(t *testing.T) {
	sw := NewSwitch("My Switch", true)

	if is, want := sw.Type, typeSwitch; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := sw.Name.GetValue(), "My Switch"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if x := sw.Name.Permissions; reflect.DeepEqual(x, characteristic.PermsRead()) == false {
		t.Fatal(x)
	}
	if x := sw.On.Permissions; reflect.DeepEqual(x, characteristic.PermsAll()) == false {
		t.Fatal(x)
	}
	if is, want := sw.On.GetValue(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
