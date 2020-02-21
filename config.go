package hc

import (
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"net"
	"reflect"

	"github.com/brutella/hc/util"
	"github.com/xiam/to"
)

// Config holds configuration options.
type Config struct {
	// Path to the storage
	// When empty, the tranport stores the data inside a folder named exactly like the accessory
	StoragePath string

	// Port on which transport is reachable e.g. 12345
	// When empty, the transport uses a random port
	Port string

	// Deprecated: Specifying a static IP is discouraged.
	IP string

	// Pin with has to be entered on iOS client to pair with the accessory
	// When empty, the pin 00102003 is used
	Pin string

	// SetupId used for setup code should be 4 uppercase letters
	SetupId string

	name         string // Accessory name
	id           string // Accessory id
	servePort    int    // Actual port the server listens at (might be differen than Port field)
	version      int64  // Accessory content version (c#)
	categoryId   uint8  // Accessory category (ci)
	state        int64  // Accessory state (s#)
	protocol     string // Protocol version, default 1.0 (pv)
	discoverable bool   // Flag if accessory is discoverable (sf)
	mfiCompliant bool   // Flag if accessory if Mfi compliant (ff)
	configHash   []byte
}

func defaultConfig(name string) *Config {
	return &Config{
		StoragePath:  name,
		Pin:          "00102003", // default pin
		Port:         "",         // empty string means that we get port from assigned by the system
		SetupId:      "HOME",     // default setup id
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
func (cfg Config) txtRecords() map[string]string {
	return map[string]string{
		"pv": cfg.protocol,
		"id": cfg.id,
		"c#": fmt.Sprintf("%d", cfg.version),
		"s#": fmt.Sprintf("%d", cfg.state),
		"sf": fmt.Sprintf("%d", to.Int64(cfg.discoverable)),
		"ff": fmt.Sprintf("%d", to.Int64(cfg.mfiCompliant)),
		"md": cfg.name,
		"ci": fmt.Sprintf("%d", cfg.categoryId),
		"sh": cfg.setupHash(),
	}
}

func (cfg *Config) setupHash() string {
	hashvalue := fmt.Sprintf("%s%s", cfg.SetupId, cfg.id)
	sum := sha512.Sum512([]byte(hashvalue))
	// use only first 4 bytes
	code := []byte{sum[0], sum[1], sum[2], sum[3]}
	encoded := base64.StdEncoding.EncodeToString(code)
	return encoded
}

func (cfg *Config) XHMURI(flag util.SetupFlag) (string, error) {
	flags := []util.SetupFlag{flag}
	return util.XHMURI(cfg.Pin, cfg.SetupId, cfg.categoryId, flags)
}

// loads load the id, version and config hash
func (cfg *Config) load(storage util.Storage) {
	if b, err := storage.Get("uuid"); err == nil && len(b) > 0 {
		cfg.id = string(b)
	}

	if b, err := storage.Get("version"); err == nil && len(b) > 0 {
		cfg.version = to.Int64(string(b))
	}

	if b, err := storage.Get("configHash"); err == nil && len(b) > 0 {
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

	if setupid := other.SetupId; len(setupid) > 0 {
		cfg.SetupId = setupid
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
