package main

import (
	"log"
	"time"

	"net/http"
	_ "net/http/pprof"
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

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	
	server.StartHttpServer()

}
