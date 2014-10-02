package netio

import(
    "github.com/brutella/hap/common"
    "io"
    "net/url"
)

type ContainerHandler interface {
    Handle(common.Container) (common.Container, error)
}

type PairVerifyHandler interface {
    ContainerHandler
    SharedKey() [32]byte
}

type AccessoriesHandler interface {
    HandleGetAccessories() (io.Reader, error)
}

type CharacteristicsHandler interface {
    HandleGetCharacteristics(url.Values) (io.Reader, error)
    HandleUpdateCharacteristics(io.Reader) error
}