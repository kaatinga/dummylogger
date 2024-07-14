package dummylogger

import (
	"sync"
)

type I interface {
	Errorf(format string, v ...interface{})
}

var logger = newDummyLogger()

var setLoggerOnce sync.Once

func Get() I {
	return logger
}

func Set(l I) {
	setLoggerOnce.Do(func() {
		logger = l
	})
}
