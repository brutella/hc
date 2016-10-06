package log

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	// Debug generates debug lines of output with a "DEBUG" prefix.
	// By default the lines are written to /dev/null.
	Debug = &Logger{log.New(ioutil.Discard, "DEBUG ", log.LstdFlags)}

	// Info generates debug lines of output with a "INFO" prefix.
	// By default the lines are written to stdout.
	Info = &Logger{log.New(os.Stdout, "INFO ", log.LstdFlags)}
)

// Logger is a wrapper for log.Logger and provides
// methods to enable and disable logging.
type Logger struct {
	*log.Logger
}

// Disable sets the logging output to /dev/null.
func (l *Logger) Disable() {
	l.SetOutput(ioutil.Discard)
}

// Enable sets the logging output to stdout.
func (l *Logger) Enable() {
	l.SetOutput(os.Stdout)
}
