package pair

import(
    "io"
)

type Handler interface {
    HandleContainer(Container) (Container, error)
    HandleReader(io.Reader) (io.Reader, error)
}
