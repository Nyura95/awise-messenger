package config

import (
	"errors"

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
var env string

// Start the configuration
func Start(environment string) {
	configuration := Config{}
	err := gonfig.GetConf("./config/config."+environment+".json", &configuration)
	if err != nil {
		panic(err)
	}
	env = environment
	config = configuration
}

// GetConfig return the configuration
func GetConfig() (Config, error) {
	if env == "" {
		return config, errors.New("The configuration is not started")
	}
	return config, nil
}

// GetEnv for know the environment of the server
func GetEnv() (string, error) {
	if env == "" {
		return env, errors.New("The configuration is not started")
	}
	return env, nil
}
