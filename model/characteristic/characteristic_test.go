package characteristic

import (
	"net"
	"testing"
)

func TestCharacteristicSetValuesOfWrongType(t *testing.T) {
	var value int = 5
	c := NewCharacteristic(value, FormatInt, CharTypePowerState, nil)

	c.SetValue(float64(20.5))

	if is, want := c.Value, 20; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.SetValue("91")

	if is, want := c.Value, 91; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	c.SetValue(true)

	if is, want := c.Value, 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCharacteristicLocalDelegate(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)

	var oldValue interface{}
	var newValue interface{}

	c.OnChange(func(c *Characteristic, new, old interface{}) {
		newValue = new
		oldValue = old
	})

	c.SetValue(10)
	c.SetValueFromConnection(20, TestConn)

	if is, want := oldValue, 5; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := newValue, 10; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCharacteristicRemoteDelegate(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)

	var oldValue interface{}
	var newValue interface{}
	c.OnConnChange(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		if conn != TestConn {
			t.Fatal(conn)
		}
		newValue = new
		oldValue = old
	})

	c.SetValueFromConnection(10, TestConn)
	c.SetValue(20)

	if is, want := oldValue, 5; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := newValue, 10; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestNoValueChange(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)

	changed := false
	c.OnConnChange(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		if conn != TestConn {
			t.Fatal(conn)
		}
		changed = true
	})

	c.OnChange(func(c *Characteristic, new, old interface{}) {
		changed = true
	})

	c.SetValue(5)
	c.SetValueFromConnection(5, TestConn)

	if changed != false {
		t.Fatal(changed)
	}
}

func TestReadOnlyValue(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, PermsRead())

	remoteChanged := false
	localChanged := false
	c.OnConnChange(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		if conn != TestConn {
			t.Fatal(conn)
		}
		remoteChanged = true
	})

	c.OnChange(func(c *Characteristic, new, old interface{}) {
		localChanged = true
	})

	c.SetValue(10)
	c.SetValueFromConnection(11, TestConn)

	if is, want := c.GetValue(), 10; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if remoteChanged != false {
		t.Fatal(remoteChanged)
	}
	if localChanged != true {
		t.Fatal(localChanged)
	}
}

func TestEqual(t *testing.T) {
	c1 := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)
	c2 := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)
	if c1.Equal(c2) == false {
		t.Fatal("characteristics not the same")
	}
}
