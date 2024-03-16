package dummylogger

import (
	"log"
	"sync"
)

type dummyLoggerI interface {
	Errorf(format string, v ...interface{})
}

var logger = newDummyLogger()

type dummyLoggerType struct{}

func newDummyLogger() dummyLoggerI {
	return dummyLogger
}

func (d *dummyLoggerType) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

var dummyLogger = &dummyLoggerType{}

var setLoggerOnce sync.Once

func Get() dummyLoggerI {
	return logger
}

func Set(l dummyLoggerI) {
	setLoggerOnce.Do(func() {
		logger = l
	})
}
