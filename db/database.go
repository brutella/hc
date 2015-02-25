package db

import (
	"github.com/brutella/hc/common"
	"github.com/gosexy/to"
)

// Database stores clients and dns persistently.
type Database interface {
	// ClientWithName returns the client referenced by name
	ClientWithName(name string) Client

	// SaveClient saves a client in the database
	SaveClient(client Client) error

	// DeleteClient deletes a client from the database
	DeleteClient(client Client)

	// DnsWithName returns the dns references by name
	DnsWithName(name string) Dns

	// SaveDns saves the dns in the database
	SaveDns(dns Dns) error

	// DeleteDns deletes a dns from the database
	DeleteDns(dns Dns)
}

type database struct {
	storage common.Storage
}

// NewDatabase returns a database which stores data into the folder specified by the argument string.
func NewDatabase(path string) (Database, error) {
	storage, err := common.NewFileStorage(path)
	if err != nil {
		return nil, err
	}

	return NewDatabaseWithStorage(storage), nil
}

// NewDatabaseWithStorage returns a database which uses the argument storage to store data.
func NewDatabaseWithStorage(storage common.Storage) Database {
	c := database{storage: storage}

	return &c
}

// ClientWithName returns a client for a specific name
// The method tries to load the ltpk from disk and returns initialized client object.
// The method returns nil when no file for this client could be found.
func (m *database) ClientWithName(name string) Client {
	data, err := m.storage.Get(keyForClientName(name))

	if len(data) > 0 && err == nil {
		client := NewClient(name, data)
		return client
	}

	return nil
}

// SaveClient stores the long-term public key of the client as {client-name}.ltpk to disk.
func (m *database) SaveClient(client Client) error {
	if len(client.PublicKey()) == 0 {
		return common.NewErrorf("No public key to save for client%s\n", client.Name())
	}

	return m.storage.Set(keyForClientName(client.Name()), client.PublicKey())
}

func (db *database) DeleteClient(client Client) {
	db.storage.Delete(keyForClientName(client.Name()))
}

func keyForClientName(name string) string {
	return name + ".ltpk"
}

func (db *database) DnsWithName(name string) Dns {
	c_data, err := db.storage.Get(configurationKeyForDnsName(name))
	s_data, err := db.storage.Get(stateKeyForDnsName(name))

	if len(c_data) > 0 && err == nil && len(s_data) > 0 {
		return NewDns(name, to.Int64(string(c_data)), to.Int64(string(s_data)))
	}

	return nil
}

func (db *database) SaveDns(dns Dns) error {
	configuration := to.String(dns.Configuration())
	state := to.String(dns.State())
	err := db.storage.Set(configurationKeyForDnsName(dns.Name()), []byte(configuration))
	if err != nil {
		return err
	}

	return db.storage.Set(stateKeyForDnsName(dns.Name()), []byte(state))
}

func (db *database) DeleteDns(dns Dns) {
	db.storage.Delete(configurationKeyForDnsName(dns.Name()))
	db.storage.Delete(stateKeyForDnsName(dns.Name()))
}

func configurationKeyForDnsName(name string) string {
	return name + ".configuration"
}

func stateKeyForDnsName(name string) string {
	return name + ".state"
}
