package util

import (
	"bytes"
	"os"
	"path"
	"path/filepath"
)

type fileStorage struct {
	dirPath string
}

// NewTempFileStorage returns a new storage inside temporary folder.
func NewTempFileStorage() (Storage, error) {
	dir := RandomHexString()
	return NewFileStorage(path.Join(os.TempDir(), dir))
}

// NewFileStorage create a file storage for the specified directory.
// The folder is created if necessary. Every key-value pair is stored in a seperate file.
func NewFileStorage(dir string) (Storage, error) {
	path, err := filepath.Abs(dir)
	if err != nil {
		return nil, err
	}

	// Why 0777?
	// Read http://unix.stackexchange.com/questions/21251/why-do-directories-need-the-executable-x-permission-to-be-opened
	err = os.MkdirAll(path, 0777)
	return &fileStorage{dirPath: path}, err
}

// Set sets the value for a specific key.
func (f *fileStorage) Set(key string, value []byte) error {
	file, err := f.fileForWrite(key)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(value)
	return err
}

// Get returns the value for a specific key.
func (f *fileStorage) Get(key string) ([]byte, error) {
	file, err := f.fileForRead(key)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var b bytes.Buffer
	var buffer = make([]byte, 32)
	for {
		n, _ := file.Read(buffer)
		if n > 0 {
			b.Write(buffer[:n])
		} else {
			break
		}
	}

	return b.Bytes(), nil
}

// Delete removes the file for the corresponding key.
func (f *fileStorage) Delete(key string) error {
	return os.Remove(f.filePathToFile(key))
}

// Private
func (f *fileStorage) dir() string {
	return f.dirPath
}

func (f *fileStorage) filePathToFile(file string) string {
	return filepath.Join(f.dir(), file)
}

func (f *fileStorage) fileForWrite(key string) (*os.File, error) {
	return os.OpenFile(f.filePathToFile(key), os.O_WRONLY|os.O_CREATE, 0666)
}

func (f *fileStorage) fileForRead(key string) (*os.File, error) {
	return os.OpenFile(f.filePathToFile(key), os.O_RDONLY, 0666)
}
