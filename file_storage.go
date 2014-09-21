package gohap

import(
    "os"
    "path/filepath"
    "bytes"
)

type fileStorage struct {
    Storage
    dir_path string
}

func NewFileStorage(dir string) Storage {
    return &fileStorage{dir_path: dir}
}

func (f *fileStorage) Update(key string, value []byte) error {
    file, err := f.fileForWrite(key)
    
    if err != nil {
        return err
    }
    
    defer file.Close()
    
    _, err = file.Write(value)
    return err
}

func (f *fileStorage) Get(key string) ([]byte, error) {
    file, err := f.fileForRead(key)
    
    if err != nil {
        return nil, err
    }
    
    defer file.Close()
    
    var b bytes.Buffer
    var buffer []byte = make([]byte, 32)
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

func (f *fileStorage) Delete(key string) error {
    return os.Remove(f.filePathToFile(key))
}


// Private

func (f *fileStorage) dir() string {
    return filepath.Dir(f.dir_path)
}

func (f *fileStorage) filePathToFile(file string) string {
    return filepath.Join([]string{f.dir(), file}...)
}

func (f *fileStorage) fileForWrite(key string) (*os.File, error) {
    return os.OpenFile(f.filePathToFile(key), os.O_WRONLY | os.O_CREATE, 0666)
}

func (f *fileStorage) fileForRead(key string) (*os.File, error) {
    return os.OpenFile(f.filePathToFile(key), os.O_RDONLY, 0666)
}