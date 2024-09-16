package config

var context *Application

type Application struct {
	ConfigProperties *Properties
}

func CreateApplication() *Application {
	properties := GetConfigProperties()
	context = &Application{ConfigProperties: &properties}
	return context
}
