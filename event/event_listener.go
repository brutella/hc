package event

// EventListener handles events
type EventListener interface {
	Handle(e interface{})
}
