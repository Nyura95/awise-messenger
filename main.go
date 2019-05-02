package main

import (
	"messenger/config"
	"messenger/models"
	"messenger/server"
	"messenger/socket"
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
