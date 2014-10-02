package common

type Storage interface {
    Set(key string, value []byte) error
    Delete(key string) error
    Get(key string) ([]byte, error)
}