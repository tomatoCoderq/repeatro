package main

import (
	"time"

	"repeatro/internal/config"
	"repeatro/internal/repositories"
	"repeatro/internal/security"
	"repeatro/internal/server"
)

func main() {
	config := config.InitConfig("config")

	security := security.Security{ExpirationDelta: 600 * time.Minute}
	security.GetKyes()

	db := repositories.InitDatabase(config)

	repositories.InitGooseMigration(db)

	server := server.InitHTTPServer(config, db, security)

	server.StartHttpServer()
}
