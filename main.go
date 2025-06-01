package main

import "repeatro/server"

func main() {
	config := InitConfig("config")

	db := server.InitDatabase(config)

	server := server.InitHTTPServer(config, db)

	server.StartHttpServer()
}