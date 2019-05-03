package main

import (
	"awise-messenger/config"
	"awise-messenger/models"
	"awise-messenger/server"
	"awise-messenger/socket"
)

func main() {
	// Instanciation de la configuration
	config.Start()
	// Instanciation du pool mysql
	models.InitDb()
	// Lancement du serveur http
	go server.Start()

	socket.Start()
}
