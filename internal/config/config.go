package config

type Properties struct {
	ServerProperties  *ServerProperties  `yaml:"server"`
	LoggingProperties *LoggingProperties `yaml:"logging"`
}

type ServerProperties struct {
	Port string `yaml:"port"`
}

type LoggingProperties struct {
	Level string `yaml:"level"`
}
