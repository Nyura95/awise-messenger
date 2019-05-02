package config

import (
	"github.com/tkanos/gonfig"
)

// Config app
type Config struct {
	User       string
	Host       string
	Password   string
	Database   string
	Port       int
	SocketPort int
}

var config Config

// Start the configuration
func Start() {
	configuration := Config{}
	err := gonfig.GetConf("./config.json", &configuration)
	if err != nil {
		panic(err)
	}
	config = configuration
}

// GetConfig return Config
func GetConfig() Config {
	return config
}
