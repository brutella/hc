package db

import (
	"encoding/hex"
	"encoding/json"
	"github.com/brutella/hc/util"
)

// Database stores entities
type Database interface {
	// EntityWithName returns the entity referenced by name
	EntityWithName(name string) (Entity, error)

	// SaveEntity saves a entity in the database
	SaveEntity(entity Entity) error

	// DeleteEntity deletes a entity from the database
	DeleteEntity(entity Entity)

	// Entities returns all entities
	Entities() ([]Entity, error)
}

type database struct {
	storage util.Storage
}

// NewTempDatabase returns a temp database
func NewTempDatabase() (Database, error) {
	storage, err := util.NewTempFileStorage()
	return NewDatabaseWithStorage(storage), err
}

// NewDatabase returns a database which stores data into the folder specified by the argument string.
func NewDatabase(path string) (Database, error) {
	storage, err := util.NewFileStorage(path)
	if err != nil {
		return nil, err
	}

	return NewDatabaseWithStorage(storage), nil
}

// NewDatabaseWithStorage returns a database which uses the argument storage to store data.
func NewDatabaseWithStorage(storage util.Storage) Database {
	c := database{storage: storage}

	return &c
}

// EntityWithName returns a entity for a specific name
// The method tries to load the ltpk from disk and returns initialized client object.
// The method returns nil when no file for this client could be found.
func (db *database) EntityWithName(name string) (e Entity, err error) {
	return db.entityForKey(toEntityKey(name))
}

// SaveEntity stores the long-term public key of the entity as {entity-name}.ltpk to disk.
func (db *database) SaveEntity(e Entity) error {
	b, err := json.Marshal(e)

	if err != nil {
		return err
	}

	return db.storage.Set(toEntityKey(e.Name), b)
}

func (db *database) DeleteEntity(e Entity) {
	db.storage.Delete(toEntityKey(e.Name))
}

func (db *database) Entities() (es []Entity, err error) {
	var e Entity
	var ks []string

	if ks, err = db.storage.KeysWithSuffix(".entity"); err == nil {
		for _, k := range ks {
			if e, err = db.entityForKey(k); err != nil {
				return nil, err
			}
			es = append(es, e)
		}
	}

	return
}

func (db *database) entityForKey(key string) (e Entity, err error) {
	var b []byte

	if b, err = db.storage.Get(key); err == nil {
		err = json.Unmarshal(b, &e)
	}

	return
}

func toEntityKey(s string) string {
	return hex.EncodeToString([]byte(s)) + ".entity"
}
