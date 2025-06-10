package main

import (
	"log"
	// "strconv"
	"repeatro/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	config := config.InitConfig("config")

	db, err := gorm.Open(postgres.Open(config.GetString("database.connection_string")))
	if err != nil {
		log.Fatalf("Error during opening database")
	}

	db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
}
