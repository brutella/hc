package common

import (
	"github.com/brutella/log"
	"time"
)

// TODO (brutella) Remove if not needed anymore
type Benchmark struct {
	name  string
	start time.Time
}

func NewBenchmark(name string) Benchmark {
	return Benchmark{name: name, start: time.Now()}
}

func (b Benchmark) Finish() {
	d := time.Since(b.start)
	log.Println(b.name, d.String())
}
