package config

import (
	"github.com/ngnhub/html_scrapper/service"
	"sync"
)

var context *AppContext
var mutex = &sync.Mutex{}

type AppContext struct {
	ConfigProperties *Properties
	Scrapper         *service.Scrapper
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
