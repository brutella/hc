package db

import(
    "github.com/brutella/hap/common"
)

type Database interface {
    ClientWithName(name string) Client
    SaveClient(client Client) error
    DeleteClient(client Client)
}

type database struct {
    storage common.Storage
}

func NewDatabase(path string) (*database, error) {
    storage, err := common.NewFileStorage(path)
    if err != nil {
        return nil, err
    }
    
    return NewDatabaseWithStorage(storage), nil
}

func NewDatabaseWithStorage(storage common.Storage) *database {
    c := database{storage: storage}
    
    return &c
}

// Returns the client for a specific name
//
// Loads the ltpk from disk and returns initialized client object
func (m *database) ClientWithName(name string) (Client) {
    data, err := m.storage.Get(keyForClientName(name))
    
    if len(data) > 0 && err == nil{
        client := NewClient(name, data)
        return client
    }
    
    return nil
}

// Stores the long-term public key of the client as {client-name}.ltpk
func (m *database) SaveClient(client Client) error {
    if len(client.PublicKey()) == 0 {
        return common.NewErrorf("No public key to save for client%s\n", client.Name())
    }
    
    return m.storage.Set(keyForClient(client), client.PublicKey())
}

// Returns they key for which the client is saved in the database
func (m *database) DeleteClient(client Client) {
    m.storage.Delete(keyForClient(client))
}

func keyForClientName(name string) string {
    return name + ".ltpk"
}

func keyForClient(client Client) string {
    return keyForClientName(client.Name())
}