// Package log implements optional logging based on simple rules.
// The provided logging methods are similar to the standard "log" package
// for simple adaptation.
package log

import (
    slog "log"
    "strings"
)

var errorPrefix   = "[ERRO]"
var warnPrefix    = "[WARN]"
var infoPrefix    = "[INFO]"
var verbosePrefix = "[VERB]"

// Set Error to false to ignore logs with "[ERRO]" prefix
var Error = true

// Set Warn to false to ignore logs with "[WARN]" prefix
var Warn = true

// Set Info to false to ignore logs with "[INFO]" prefix
var Info = true

// Set Verbose to false to ignore logs with "[VERB]" prefix
var Verbose = true

// Println logs using log.Println
// When the first argument is prefixed with a disabled prefix, this method does nothing.
func Println(v ...interface{}) {
    shouldLog := true
    if len(v) > 0{
        if first, ok := v[0].(string); ok == true {
            if strings.HasPrefix(first, errorPrefix) == true {
                shouldLog = Error
            } else if strings.HasPrefix(first, warnPrefix) == true {
                shouldLog = Warn
            } else if strings.HasPrefix(first, infoPrefix) == true {
                shouldLog = Info
            } else if strings.HasPrefix(first, verbosePrefix) == true {
                shouldLog = Verbose
            }
        }
    }
    
    if shouldLog == true {
        slog.Println(v...)
    }
}

// Printf logs using log.Printf
// When format string is prefixed with a disabled prefix, this method does nothing.
func Printf(format string, v ...interface{}) {
    shouldLog := true
    
    if strings.HasPrefix(format, errorPrefix) == true {
        shouldLog = Error
    } else if strings.HasPrefix(format, warnPrefix) == true {
        shouldLog = Warn
    } else if strings.HasPrefix(format, infoPrefix) == true {
        shouldLog = Info
    } else if strings.HasPrefix(format, verbosePrefix) == true {
        shouldLog = Verbose
    }
    
    if shouldLog == true {
        slog.Printf(format, v...)
    }
}

// Fatalf logs using log.Fatalf
func Fatalf(format string, v ...interface{}) {
    slog.Fatalf(format, v...)
}

// Fatalln logs using log.Fatalln
func Fatalln(v ...interface{}) {
    slog.Fatalln(v...)
}

// Fatal logs using log.Fatal
func Fatal(v ...interface{}) {
    slog.Fatal(v...)
}