package tlv8

import (
	"fmt"
	"reflect"
	"testing"
)

type user struct {
	Name    string    `tlv8:"1"`
	Type    uint8     `tlv8:"2"`
	Pwd     *password `tlv8:"3"`
	Enabled bool      `tlv8:"4"`
	Id      uint16    `tlv8:"5"`
	Uid     uint32    `tlv8:"6"`
}

func (u user) String() string {
	return fmt.Sprintf("Name: %v, Email: %v, Pwd: %v", u.Name, u.Type, u.Pwd)
}

type password struct {
	Plaintext string `tlv8:"4"`
	Hash      byte   `tlv8:"5"`
}

func (p password) String() string {
	return fmt.Sprintf("Plaintext: %s, Hash: %v", p.Plaintext, p.Hash)
}

var pwd = password{"asdf", 8}
var u = user{"Matthias", 4, &pwd, true, 400, 40100}

func TestMarshal(t *testing.T) {
	tlv8, err := Marshal(u)
	if err != nil {
		t.Fatal(err)
	}

	expect := []byte{
		1, 8, 77, 97, 116, 116, 104, 105, 97, 115,
		2, 1, 4,
		3, 9,
		4, 4, 97, 115, 100, 102,
		5, 1, 8,
		4, 1, 1,
		5, 2, 144, 1,
		6, 4, 164, 156, 0, 0,
	}
	if is, want := tlv8, expect; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestMarshalList(t *testing.T) {
	type Object struct {
		Id byte `tlv8:"1"`
	}
	type List struct {
		Elements []Object `tlv8:"-"`
	}

	objs := []Object{Object{1}, Object{2}}
	l := List{objs}

	tlv8, _ := Marshal(l)
	expect := []byte{
		1, 1, 1,
		0, 0,
		1, 1, 2,
	}
	if is, want := tlv8, expect; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestUnmarshalList(t *testing.T) {
	type Object struct {
		Id byte `tlv8:"1"`
	}

	objs := []Object{Object{1}, Object{2}}
	tlv8, err := Marshal(objs)
	if err != nil {
		t.Fatal(err)
	}

	var objects []Object
	err = Unmarshal(tlv8, &objects)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnmarshalUint16(t *testing.T) {
	type Object struct {
		Id uint16 `tlv8:"1"`
	}

	tlv8, _ := Marshal(Object{1000})

	var obj Object
	err := Unmarshal(tlv8, &obj)
	if err != nil {
		t.Fatal(err)
	}

	if x := obj.Id; x != 1000 {
		t.Fatal(x)
	}
}

func TestUnmarshal(t *testing.T) {
	tlv8, _ := Marshal(u)

	var usr user
	err := Unmarshal(tlv8, &usr)
	if err != nil {
		t.Fatal(err)
	}

	if is, want := usr.Name, "Matthias"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := usr.Type, uint8(4); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := usr.Id, uint16(400); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := usr.Uid, uint32(40100); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := usr.Enabled, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := usr.Pwd.Plaintext, "asdf"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := usr.Pwd.Hash, byte(8); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

type Attribute struct {
	Id uint8 `tlv8:"1"`
}
type Object struct {
	Value   string      `tlv8:"0"`
	Ptr     *Attribute  `tlv8:"1"`
	Byte    uint8       `tlv8:"5"`
	Bytes   []byte      `tlv8:"6"`
	Uint16  uint16      `tlv8:"7"`
	Uint32  uint32      `tlv8:"8"`
	Int16   int16       `tlv8:"9"`
	Int32   int32       `tlv8:"19"`
	Float32 float32     `tlv8:"20"`
	Struct  Attribute   `tlv8:"11"`
	Structs []Attribute `tlv8:"12"`
}

func TestMarshalObject(t *testing.T) {
	obj := Object{
		Value:   "string",
		Ptr:     &Attribute{10},
		Byte:    1,
		Bytes:   []byte{1, 2, 3},
		Uint16:  400,
		Uint32:  70000,
		Int16:   -400,
		Int32:   -70000,
		Float32: 1.234567,
		Struct:  Attribute{1},
		Structs: []Attribute{Attribute{2}, Attribute{3}},
	}
	tlv8, err := Marshal(obj)

	if err != nil {
		t.Fatal(err)
	}

	var other Object
	err = Unmarshal(tlv8, &other)
	if err != nil {
		t.Fatal(err)
	}

	if is, want := other.Value, "string"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := other.Ptr.Id, uint8(10); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := obj.Byte, uint8(1); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := obj.Bytes, []byte{1, 2, 3}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := obj.Uint16, uint16(400); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := obj.Uint32, uint32(70000); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := obj.Int16, int16(-400); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := obj.Int32, int32(-70000); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := obj.Float32, float32(1.234567); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := obj.Struct.Id, uint8(1); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := obj.Structs, []Attribute{Attribute{2}, Attribute{3}}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

type Element struct {
	Value byte `tlv8:"1"`
}

type List struct {
	Named   []Element `tlv8:"10"`
	Unnamed []Element `tlv8:"-"`
}

func TestList(t *testing.T) {
	list := List{
		Named:   []Element{Element{2}, Element{3}},
		Unnamed: []Element{Element{4}, Element{5}},
	}
	tlv8, err := Marshal(list)

	expected := []byte{
		10, 3, 1, 1, 2,
		0, 0,
		10, 3, 1, 1, 3,
		1, 1, 4, 0, 0, 1, 1, 5,
	}

	if is, want := tlv8, expected; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if err != nil {
		t.Fatal(err)
	}

	var other List
	err = Unmarshal(tlv8, &other)
	if err != nil {
		t.Fatal(err)
	}

	if is, want := other.Named, []Element{Element{2}, Element{3}}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := other.Unnamed, []Element{Element{4}, Element{5}}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
