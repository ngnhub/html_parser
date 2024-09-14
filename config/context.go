package config

import (
	"sync"
)

var context *AppContext
var mutex = &sync.Mutex{}

type AppContext struct {
	ConfigProperties *Properties
}

func GetAppContext() *AppContext {
	mutex.Lock()
	if context == nil {
		properties := GetConfigProperties()
		context = &AppContext{ConfigProperties: &properties}
	}
	mutex.Unlock()
	return context
}
