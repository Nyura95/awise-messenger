package main

import (
	"awise-messenger/config"
	"awise-messenger/models"
	"awise-messenger/server"
	"awise-messenger/socket"
	"awise-messenger/static"
)

func main() {
	// Init of the config
	config.Start("dev")
	configuration, _ := config.GetConfig()
	// Init of the pool mysql
	models.InitDb(configuration.User, configuration.Password, configuration.Host, configuration.Database)
	defer models.Close()

	// Launch of the http server
	go server.Start()

	// Launch of the static server for the front
	go static.Start()

	// Launch of the socket server
	socket.Start()

}
