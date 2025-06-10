package main

import (
	"repeatro/internal/config"
	"repeatro/internal/repositories"
	"repeatro/internal/security"
	"repeatro/internal/server"
	"repeatro/migrations"

	"github.com/pressly/goose/v3"
)

func main() {
	config := config.InitConfig("config")

	security := security.Security{}
	security.GenerateKey()

	db := repositories.InitDatabase(config)

	goose.SetBaseFS(migrations.EmbedMigrations)

	// Automatically fetch migrations
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	trueDB, err := db.DB()
	if err != nil {
		panic(err)
	}

    if err := goose.Up(trueDB, "."); err != nil {
        panic(err)
    }

	server := server.InitHTTPServer(config, db, security)

	server.StartHttpServer()
}