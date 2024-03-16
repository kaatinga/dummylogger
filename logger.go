package dummylogger

import (
	"log"
	"sync"
)

type I interface {
	Errorf(format string, v ...interface{})
}

var logger = newDummyLogger()

type dummyLoggerType struct{}

func newDummyLogger() I {
	return dummyLogger
}

func (d *dummyLoggerType) Errorf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

var dummyLogger = &dummyLoggerType{}

var setLoggerOnce sync.Once

func Get() I {
	return logger
}

func Set(l I) {
	setLoggerOnce.Do(func() {
		logger = l
	})
}
