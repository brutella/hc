package db

import(
    "github.com/brutella/hap"
    "github.com/brutella/hap/common"
)

type Manager struct {
    storage hap.Storage
}

func NewManager(storage hap.Storage) *Manager {
    c := Manager{storage: storage}
    
    return &c
}

// Returns the client for a specific name
//
// Loads the ltpk from disk and returns initialized client object
func (m *Manager) ClientForName(name string) (*Client) {
    data, err := m.storage.Get(name + ".ltpk")
    
    if len(data) > 0 && err == nil{
        client := NewClient(name, data)
        return client
    }
    
    return nil
}

// Stores the long-term public key of the client as {client-name}.ltpk
func (m *Manager) SaveClient(client *Client) error {
    if len(client.PublicKey) == 0 {
        return common.NewErrorf("No public key to save for client%s\n", client.Name)
    }
    
    return m.storage.Set(client.Name + ".ltpk", client.PublicKey)
}

func (m *Manager) DeleteClient(client *Client) {
    m.storage.Delete(client.Name + ".ltpk")
}


// Returns the party for a specific name
//
// Loads the ltpk, ltsk and serial number from disk
func (m *Manager) PartyWithName(name string) (*Party) {
    serial, _ := m.storage.Get(name + ".serial")
    ltpk, _ := m.storage.Get(name + ".ltpk")
    ltsk, _ := m.storage.Get(name + ".ltsk")
    
    return NewParty(name, string(serial), ltpk, ltsk)
}

// Stores the long-term public key of the party as {client-name}.ltpk
// Stores the long-term secrekt key of the party as {client-name}.ltsk
// Stores the serial number of the party as {client-name}.serial
func (m *Manager) SaveParty(p *Party) error {
    if len(p.PublicKey) > 0 {
        err := m.storage.Set(p.Name + ".ltpk", p.PublicKey)
        if err != nil {
            return err
        }
    } else {
        m.storage.Delete(p.Name + ".ltpk")
    }
    
    if len(p.SerialNumber) > 0 {
        err := m.storage.Set(p.Name + ".serial", []byte(p.SerialNumber))
        if err != nil {
            return err
        }
    } else {
        m.storage.Delete(p.Name + ".serial")
    }
    
    if len(p.SecretKey) > 0 {
        err := m.storage.Set(p.Name + ".ltsk", p.SecretKey)
        if err != nil {
            return err
        }
    } else {
        m.storage.Delete(p.Name + ".ltsk")
    }
    
    return nil
}

func (m *Manager) DeleteParty(p *Party) {
    m.storage.Delete(p.Name + ".serial")
    m.storage.Delete(p.Name + ".ltpk")
    m.storage.Delete(p.Name + ".ltsk")
}