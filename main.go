package main

import (
	"awise-messenger/config"
	"awise-messenger/modelsv2"
	"awise-messenger/server"
	"awise-messenger/socketv2"
)

func main() {
	// Instanciation de la configuration
	config.Start("dev")
	configuration, _ := config.GetConfig()
	// Instanciation du pool mysql
	modelsv2.InitDb(configuration.User, configuration.Password, configuration.Host, configuration.Database)
	defer modelsv2.Close()

	// Lancement du serveur http
	go server.Start()

	// socket.Start()
	socketv2.Start()
}
