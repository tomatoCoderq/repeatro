package main

import (
	"repeatro/server"
	"repeatro/config"
)

func main() {
	config := config.InitConfig("config")

	db := server.InitDatabase(config)

	server := server.InitHTTPServer(config, db)

	server.StartHttpServer()
}