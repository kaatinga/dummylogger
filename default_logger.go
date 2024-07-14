package dummylogger

import "fmt"

type dummyLoggerType struct{}

func newDummyLogger() I {
	return dummyLogger
}

func (d *dummyLoggerType) Errorf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
}

var dummyLogger = &dummyLoggerType{}
