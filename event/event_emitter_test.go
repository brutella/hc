package event

import (
	"testing"
)

type testListener struct {
	last interface{}
}

func (l *testListener) Handle(e interface{}) {
	l.last = e
}

func TestEmitter(t *testing.T) {
	e := NewEmitter()

	l := &testListener{}
	e.AddListener(l)

	e.Emit(10)

	if x := l.last; x != 10 {
		t.Fatal(x)
	}
}
