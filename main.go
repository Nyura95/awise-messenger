package main

import (
	"awise-messenger/config"
	"awise-messenger/models"
	"awise-messenger/server"
	socketv2 "awise-messenger/socketV2"
)

func main() {
	// Instanciation de la configuration
	config.Start()
	// Instanciation du pool mysql
	models.InitDb()
	// Lancement du serveur http
	go server.Start()

	// socket.Start()
	socketv2.Start()
}
