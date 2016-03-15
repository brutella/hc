package hap

import (
	"reflect"
	"testing"
)

func TestMDNS(t *testing.T) {
	mdns := NewMDNSService("My MDNS Service", "1234", "127.0.0.1", 5010, 1)
	expect := []string{
		"pv=1.0",
		"id=1234",
		"c#=1",
		"s#=1",
		"sf=1",
		"ff=0",
		"md=My MDNS Service",
		"ci=1",
	}
	if x := mdns.txtRecords(); reflect.DeepEqual(x, expect) == false {
		t.Fatal(expect)
	}
}

func TestReachable(t *testing.T) {
	mdns := NewMDNSService("My MDNS Service", "1234", "127.0.0.1", 5010, 1)
	expect := []string{
		"pv=1.0",
		"id=1234",
		"c#=1",
		"s#=1",
		"sf=0",
		"ff=0",
		"md=My MDNS Service",
		"ci=1",
	}
	mdns.SetReachable(false)

	if x := mdns.txtRecords(); reflect.DeepEqual(x, expect) == false {
		t.Fatal(expect)
	}
}
