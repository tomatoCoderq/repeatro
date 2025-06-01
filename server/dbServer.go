package server

import (
	"log"
	// "strconv"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config *viper.Viper) *gorm.DB {
	dsn := "host=localhost port=5432 user=tomatocoder password=postgres dbname=repeatro sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalf("Error during opening database")
	}
	
	return db


	// connectionString := config.GetString("database.connection_string")
	// max_idle_connections := config.GetString("database.max_idle_connections")
	// max_open_connections := config.GetString("database.max_open_connections")
	// connection_max_lifetime := config.GetString("database.conecction_max_lifetime")
	// driver_name := config.GetString("database.driver_name")
	// if connectionString == "" {
	// 	log.Fatalf("Database connection string is missing")
	// }
	// dbHandler, err := sql.Open(driver_name, connectionString)
	// if err != nil {
	// 	log.Fatalf("Error %v", err)
	// }
	// max_idle_connections_int, err := strconv.Atoi(max_idle_connections)
	// if err != nil {
	// 	log.Fatalf("Error during conversion")
	// }

	// max_open_connections_int, err := strconv.Atoi(max_open_connections)
	// if err != nil {
	// 	log.Fatalf("Error during conversion")
	// }

	// connection_max_lifetime_duration, err := str2duration.ParseDuration(connection_max_lifetime)
	// if err != nil {
	// 	log.Fatalf("Error during conversion")
	// }

	// dbHandler.SetMaxIdleConns(max_idle_connections_int)
	// dbHandler.SetMaxOpenConns(max_open_connections_int)
	// dbHandler.SetConnMaxLifetime(connection_max_lifetime_duration)

	// err = dbHandler.Ping()
	// if err != nil {
	// 	dbHandler.Close()
	// 	log.Fatalf("Error while validatin base %v", err)
	// }
	// return dbHandler
}
