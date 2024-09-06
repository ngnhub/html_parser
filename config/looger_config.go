package config

import log "github.com/sirupsen/logrus"

func (a *AppContext) ConfigLogger() {
	properties := a.ConfigProperties.LoggingProperties
	if properties == nil {
		log.SetLevel(log.InfoLevel)
		log.Info("Logging level set to INFO")
		return
	}
	logLevel := properties.Level
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}
	if &level == nil {
		log.SetLevel(log.InfoLevel)
		log.Info("Logging level set to INFO")
	} else {
		log.SetLevel(level)
		log.Infof("Logging level set to %s", logLevel)
	}
}
