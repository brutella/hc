package common

// Storage stores bytes for a key.
type Storage interface {
	// Set sets bytes for a key
	Set(key string, value []byte) error

	// Delete removes bytes for a key
	Delete(key string) error

	// Get returns bytes for a key
	Get(key string) ([]byte, error)
}

var StdStorage Storage
