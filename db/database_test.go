package db

import (
	"os"
	"testing"
    "reflect"
)

func TestLoadUndefinedEntity(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
    
	if x := db.EntityWithName("My Name"); x != nil {
	    t.Fatal(x)
	}
}

func TestLoadEntity(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	db.SaveEntity(NewEntity("My Name", []byte{0x01}, []byte{0x02}))
	entity := db.EntityWithName("My Name")
	
    if entity == nil {
	    t.Fatal("entity not found")
	}
    
    if x := entity.PublicKey(); reflect.DeepEqual(x, []byte{0x01}) == false {
        t.Fatal(x)
    }
    
    if x := entity.PrivateKey(); reflect.DeepEqual(x, []byte{0x02}) == false {
        t.Fatal(x)
    }
}

func TestLoadEntityWithPublicKeyOnly(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	db.SaveEntity(NewEntity("Entity", []byte{0x03}, nil))
	entity := db.EntityWithName("Entity")
    if entity == nil {
	    t.Fatal("entity not found")
	}
    
    if x := entity.PublicKey(); reflect.DeepEqual(x, []byte{0x03}) == false {
        t.Fatal(x)
    }
    
    if x := entity.PrivateKey(); x != nil {
        t.Fatal(x)
    }
}

func TestDeleteEntity(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	c := NewEntity("My Name", []byte{0x01}, nil)
	db.SaveEntity(c)
	db.DeleteEntity(c)
	if x := db.EntityWithName("My Name"); x != nil {
	    t.Fatal(x)
	}
}

func TestLoadDNS(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	dns := NewDNS("My Name", 10, 20)
	db.SaveDNS(dns)
	dns = db.DNSWithName("My Name")
    
    if dns == nil {
        t.Fatal("dns not found")
    }
    
    if x := dns.Configuration(); x != 10 {
        t.Fatal(x)
    }
    
    if x := dns.State(); x != 20 {
        t.Fatal(x)
    }
}

func TestDeleteDNS(t *testing.T) {
	db, _ := NewDatabase(os.TempDir())
	dns := NewDNS("My Name", 10, 20)
	db.SaveDNS(dns)
	db.DeleteDNS(dns)
	if x := db.DNSWithName("My Name"); x != nil {
	    t.Fatal(x)
	}
}
