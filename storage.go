package gohap

type Storage interface {
    Update(key string, value []byte) error
    Delete(key string) error
    Get(key string) ([]byte, error)
}