package hc

import (
	"errors"
	"fmt"
	"log"
	"net"
	"reflect"

	"github.com/brutella/hc/util"
	"github.com/gosexy/to"
)

// Config provides  basic cfguration for an IP transport
type Config struct {
	// Path to the storage
	// When empty, the tranport stores the data inside a folder named exactly like the accessory
	StoragePath string

	// Port on which transport is reachable e.g. 12345
	// When empty, the transport uses a random port
	Port string

	// IP on which clients can connect.
	IP string

	// Pin with has to be entered on iOS client to pair with the accessory
	// When empty, the pin 00102003 is used
	Pin string

	name         string // Accessory name
	id           string // Accessory id
	servePort    int    // Actual port the server listens at (might be differen than Port field)
	version      int64  // Accessory content version (c#)
	categoryId   int    // Accessory category (ci)
	state        int64  // Accessory state (s#)
	protocol     string // Protocol version, default 1.0 (pv)
	discoverable bool   // Flag if accessory is discoverable (sf)
	mfiCompliant bool   // Flag if accessory if Mfi compliant (ff)
	configHash   []byte
}

func defaultConfig(name string) *Config {
	ip, err := getFirstLocalIPAddr()
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		StoragePath:  name,
		Pin:          "00102003", // default pin
		Port:         "",         // empty string means that we get port from assigned by the system
		IP:           ip.String(),
		name:         name,
		id:           util.MAC48Address(util.RandomHexString()),
		version:      1,
		state:        1,
		protocol:     "1.0",
		discoverable: true,
		mfiCompliant: false,
	}
}

// txtRecords returns the config formatted as mDNS txt records
func (cfg Config) txtRecords() []string {
	return []string{
		fmt.Sprintf("pv=%s", cfg.protocol),
		fmt.Sprintf("id=%s", cfg.id),
		fmt.Sprintf("c#=%d", cfg.version),
		fmt.Sprintf("s#=%d", cfg.state),
		fmt.Sprintf("sf=%d", to.Int64(cfg.discoverable)),
		fmt.Sprintf("ff=%d", to.Int64(cfg.mfiCompliant)),
		fmt.Sprintf("md=%s", cfg.name),
		fmt.Sprintf("ci=%d", cfg.categoryId),
	}
}

// loads load the id, version and config hash
func (cfg *Config) load(storage util.Storage) {
	if b, err := storage.Get("uuid"); err == nil {
		cfg.id = string(b)
	}

	if b, err := storage.Get("version"); err == nil {
		cfg.version = to.Int64(string(b))
	}

	if b, err := storage.Get("configHash"); err == nil {
		cfg.configHash = b
	}
}

// save stores the id, version and config
func (cfg *Config) save(storage util.Storage) {
	storage.Set("uuid", []byte(cfg.id))
	storage.Set("version", []byte(fmt.Sprintf("%d", cfg.version)))
	storage.Set("configHash", []byte(cfg.configHash))
}

// merge updates the StoragePath, Pin, Port and IP fields of the receiver from other.
func (cfg *Config) merge(other Config) {
	if dir := other.StoragePath; len(dir) > 0 {
		cfg.StoragePath = dir
	}

	if pin := other.Pin; len(pin) > 0 {
		cfg.Pin = pin
	}

	if port := other.Port; len(port) > 0 {
		cfg.Port = ":" + port
	}

	if ip := other.IP; len(ip) > 0 {
		cfg.IP = ip
	}
}

// updateConfigHash updates configHash of the receiver and increments version
// if new hash is different than old one.
func (cfg *Config) updateConfigHash(hash []byte) {
	if cfg.configHash != nil && reflect.DeepEqual(hash, cfg.configHash) == false {
		cfg.version += 1
	}

	cfg.configHash = hash
}

// getFirstLocalIPAddr returns the first available IP address of the local machine
// This is a fix for Beaglebone Black where net.LookupIP(hostname) return no IP address.
func getFirstLocalIPAddr() (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip == nil || ip.IsLoopback() || ip.IsUnspecified() {
			continue
		}
		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}
		return ip, nil
	}

	return nil, errors.New("Could not determine ip address")
}
