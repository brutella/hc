package db

import (
	"fmt"
	"github.com/brutella/hc/common"
	"github.com/gosexy/to"
)

// Database stores entities and dns persistently.
type Database interface {
	// EntityWithName returns the entity referenced by name
	EntityWithName(name string) Entity

	// SaveEntity saves a entity in the database
	SaveEntity(entity Entity) error

	// DeleteEntity deletes a entity from the database
	DeleteEntity(entity Entity)

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

// NewTempDatabase returns a temp database
func NewTempDatabase() (Database, error) {
	storage, err := common.NewTempFileStorage()
	return NewDatabaseWithStorage(storage), err
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

// EntityWithName returns a entity for a specific name
// The method tries to load the ltpk from disk and returns initialized client object.
// The method returns nil when no file for this client could be found.
func (m *database) EntityWithName(name string) Entity {
	publicKeyFile := publicKeyFileForEntityName(name)
	privateKeyFile := privateKeyFileForEntityName(name)

	public, err_public := m.storage.Get(publicKeyFile)
	// Ignore error for private key which is optional
	private, _ := m.storage.Get(privateKeyFile)

	if len(public) > 0 && err_public == nil {
		return NewEntity(name, public, private)
	}

	return nil
}

// SaveEntity stores the long-term public key of the entity as {entity-name}.ltpk to disk.
func (m *database) SaveEntity(entity Entity) error {
	name := entity.Name()
	if len(entity.PublicKey()) == 0 {
		return fmt.Errorf("No public key to save for entity%s\n", name)
	}

	publicKeyFile := publicKeyFileForEntityName(name)
	err := m.storage.Set(publicKeyFile, entity.PublicKey())
	if err != nil {
		return err
	}

	privateKeyFile := privateKeyFileForEntityName(name)
	return m.storage.Set(privateKeyFile, entity.PrivateKey())
}

func (db *database) DeleteEntity(entity Entity) {
	db.storage.Delete(privateKeyFileForEntityName(entity.Name()))
	db.storage.Delete(publicKeyFileForEntityName(entity.Name()))
}

func privateKeyFileForEntityName(name string) string {
	return name + ".privateKey"
}

func publicKeyFileForEntityName(name string) string {
	return name + ".publicKey"
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
