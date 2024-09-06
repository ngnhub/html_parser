package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

func GetConfigProperties() Properties {
	data, err := os.ReadFile("application.yml")
	if err != nil {
		log.Fatalf("Error reading application.yml")
	}
	configProps := Properties{}
	err = yaml.Unmarshal(data, &configProps)
	if err != nil {
		log.Fatalf("Error parsing application.yml")
	}
	return configProps
}
