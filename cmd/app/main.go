package main

import (
	"repeatro/internal/server"
	"repeatro/internal/config"
	"repeatro/internal/repositories"
)

func main() {
	config := config.InitConfig("config")

	db := repositories.InitDatabase(config)

	server := server.InitHTTPServer(config, db)

	server.StartHttpServer()
}