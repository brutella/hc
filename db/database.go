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

	// DNSWithName returns the dns references by name
	DNSWithName(name string) DNS

	// SaveDNS saves the dns in the database
	SaveDNS(dns DNS) error

	// DeleteDNS deletes a dns from the database
	DeleteDNS(dns DNS)
}

// StdDatabase is the standard database
var StdDatabase Database

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
func (db *database) EntityWithName(name string) Entity {
	publicKeyFile := publicKeyFileForEntityName(name)
	privateKeyFile := privateKeyFileForEntityName(name)

	public, err := db.storage.Get(publicKeyFile)
	// Ignore error for private key which is optional
	private, _ := db.storage.Get(privateKeyFile)

	if len(public) > 0 && err == nil {
		return NewEntity(name, public, private)
	}

	return nil
}

// SaveEntity stores the long-term public key of the entity as {entity-name}.ltpk to disk.
func (db *database) SaveEntity(entity Entity) error {
	name := entity.Name()
	if len(entity.PublicKey()) == 0 {
		return fmt.Errorf("No public key to save for entity%s\n", name)
	}

	publicKeyFile := publicKeyFileForEntityName(name)
	err := db.storage.Set(publicKeyFile, entity.PublicKey())
	if err != nil {
		return err
	}

	privateKeyFile := privateKeyFileForEntityName(name)
	return db.storage.Set(privateKeyFile, entity.PrivateKey())
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

func (db *database) DNSWithName(name string) DNS {
	config, err := db.storage.Get(configurationKeyForDNSName(name))
	state, err := db.storage.Get(stateKeyForDNSName(name))

	if len(config) > 0 && err == nil && len(state) > 0 {
		return NewDNS(name, to.Int64(string(config)), to.Int64(string(state)))
	}

	return nil
}

func (db *database) SaveDNS(dns DNS) error {
	config := to.String(dns.Configuration())
	state := to.String(dns.State())
	err := db.storage.Set(configurationKeyForDNSName(dns.Name()), []byte(config))
	if err != nil {
		return err
	}

	return db.storage.Set(stateKeyForDNSName(dns.Name()), []byte(state))
}

func (db *database) DeleteDNS(dns DNS) {
	db.storage.Delete(configurationKeyForDNSName(dns.Name()))
	db.storage.Delete(stateKeyForDNSName(dns.Name()))
}

func configurationKeyForDNSName(name string) string {
	return name + ".configuration"
}

func stateKeyForDNSName(name string) string {
	return name + ".state"
}
