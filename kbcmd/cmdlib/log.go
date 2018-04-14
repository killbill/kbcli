package cmdlib

import (
	"fmt"
	"log"
)

type logger struct{}

// NewLogger returns new logger.
func NewLogger() Logger {
	return &logger{}
}

// Infof - print log
func (*logger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("I: " + msg)
}

// Warningf - print warning message
func (*logger) Warningf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("W: " + msg)
}

// Errorf - print error
func (*logger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("E: " + msg)
}
