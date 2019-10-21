package main

import (
	"awise-messenger/config"
	"awise-messenger/models"
	"awise-messenger/server"
	"awise-messenger/socket"
)

func main() {
	// Instanciation de la configuration
	config.Start("dev")
	configuration, _ := config.GetConfig()
	// Instanciation du pool mysql
	models.InitDb(configuration.User, configuration.Password, configuration.Host, configuration.Database)
	defer models.Close()

	// Lancement du serveur http
	go server.Start()

	// socket.Start()
	socket.Start()
}
