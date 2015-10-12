package event

// Emitter emits events to listeners
type Emitter interface {

	// Emit emits the event to all listeners
	Emit(ev interface{})

	// AddListener adds a listener to the event stream
	AddListener(l EventListener)
}

type eventEmitter struct {
	ls []EventListener
}

// NewEmitter returns a new event emitter
func NewEmitter() Emitter {
	return &eventEmitter{
		ls: make([]EventListener, 0),
	}
}

func (e *eventEmitter) Emit(ev interface{}) {
	for _, l := range e.ls {
		l.Handle(ev)
	}
}

func (e *eventEmitter) AddListener(l EventListener) {
	e.ls = append(e.ls, l)
}
