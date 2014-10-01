package pair

type PairingHandler interface {
    Handle(Container) (Container, error)
}
