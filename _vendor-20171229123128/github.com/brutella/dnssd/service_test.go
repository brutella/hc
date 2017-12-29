package dnssd

import (
	"net"
	"testing"
)

func TestParseServiceInstanceName(t *testing.T) {
	instance, service, domain := parseServiceInstanceName("Test._hap._tcp.local.")

	if is, want := instance, "Test"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := service, "_hap._tcp."; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := domain, "local."; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestParseHostname(t *testing.T) {
	name, domain := parseHostname("Computer.local.")

	if name != "Computer" {
		t.Fatalf("%s != Computer", name)
	}

	if is, want := domain, "local."; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestParseHostnameTrailingDomain(t *testing.T) {
	name, domain := parseHostname("Computer.local")

	if name != "Computer" {
		t.Fatalf("%s != Computer", name)
	}

	if is, want := domain, "local."; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestParseHostnameWithoutDomain(t *testing.T) {
	name, domain := parseHostname("Computer.")

	if is, want := name, "Computer"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := domain, ""; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestParseHostnameWithoutTrailingDot(t *testing.T) {
	name, domain := parseHostname("Computer")

	if is, want := name, "Computer"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := domain, ""; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestNewServiceWithoutHostname(t *testing.T) {
	i := NewService("Test", "_asdf._tcp", "local.", "", nil, 1234)

	if len(i.IPs) == 0 {
		t.Fatal("Expected ips")
	}

	if len(i.IfaceIPs) == 0 {
		t.Fatal("Expected interface ips")
	}
}

func TestNewServiceWithoutIP(t *testing.T) {
	i := NewService("Test", "_asdf._tcp", "local.", "Computer", nil, 1234)

	if len(i.IPs) == 0 {
		t.Fatal("Expected ips")
	}

	if len(i.IfaceIPs) == 0 {
		t.Fatal("Expected interface ips")
	}
}

func TestNewServiceIP(t *testing.T) {
	i := NewService("Test", "_asdf._tcp", "local.", "Computer", []net.IP{net.ParseIP("127.0.0.1")}, 1234)

	if is, want := len(i.IPs), 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if x := i.IfaceIPs; x != nil {
		t.Fatal(x)
	}
}

func TestSanitizeHostname(t *testing.T) {
	host := sanitizeHostname("My Computer")
	if is, want := host, "My-Computer"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
