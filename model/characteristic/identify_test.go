package characteristic

import (
	"net"
	"testing"
)

func TestWriteOnlyIdentifyCharacteristic(t *testing.T) {
	i := NewIdentify()

	if is, want := i.Type, CharTypeIdentify; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if x := i.GetValue(); x != nil {
		t.Fatal(x)
	}

	i.SetBool(true)

	if x := i.GetValue(); x != nil {
		t.Fatal(x)
	}

	i.SetValueFromConnection(true, TestConn)

	if x := i.GetValue(); x != nil {
		t.Fatal(x)
	}

	i.SetValue(true)

	if x := i.GetValue(); x != nil {
		t.Fatal(x)
	}
}

func TestWriteOnlyCharacteristicRemoteDelegate(t *testing.T) {
	i := NewIdentify()

	var oldValue interface{}
	var newValue interface{}
	i.OnConnChange(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		newValue = new
		oldValue = old
	})

	i.SetValueFromConnection(true, TestConn)
	if oldValue != nil {
		t.Fatal(oldValue)
	}

	if newValue != true {
		t.Fatal(newValue)
	}

	if x := i.GetValue(); x != nil {
		t.Fatal(x)
	}

	i.SetValueFromConnection(false, TestConn)

	if oldValue != nil {
		t.Fatal(oldValue)
	}

	if newValue != false {
		t.Fatal(newValue)
	}

	if x := i.GetValue(); x != nil {
		t.Fatal(x)
	}
}
