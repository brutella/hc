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

	if is, want := service, "_hap._tcp"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := domain, "local"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestParseHostname(t *testing.T) {
	name, domain := parseHostname("Computer.local.")

	if name != "Computer" {
		t.Fatalf("%s != Computer", name)
	}

	if is, want := domain, "local"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestParseHostnameTrailingDomain(t *testing.T) {
	name, domain := parseHostname("Computer.local")

	if name != "Computer" {
		t.Fatalf("%s != Computer", name)
	}

	if is, want := domain, "local"; is != want {
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

func TestNewServiceWithMinimalConfig(t *testing.T) {
	cfg := Config{
		Name: "Test",
		Type: "_asdf._tcp",
		Port: 1234,
	}

	sv, err := NewService(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if len(sv.Host) == 0 {
		t.Fatal("Expected hostname")
	}

	if is, want := sv.Domain, "local"; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if is, want := len(sv.IPs), 0; is != want {
		t.Fatalf("%v != %v", is, want)
	}

	if len(sv.IfaceIPs) == 0 {
		t.Fatal("Expected interface ips")
	}
}

func TestNewServiceWithExplicitIP(t *testing.T) {
	cfg := Config{
		Name: "Test",
		Type: "_asdf._tcp",
		IPs:  []net.IP{net.ParseIP("127.0.0.1")},
		Port: 1234,
	}
	sv, err := NewService(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if is, want := len(sv.IPs), 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := sv.IPs[0].String(), "127.0.0.1"; is != want {
		t.Fatalf("%v != %v", is, want)
	}
}

func TestSanitizeHostname(t *testing.T) {
	host := sanitizeHostname("My Computer")
	if is, want := host, "My-Computer"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
