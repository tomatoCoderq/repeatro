package main

import (
	"log/slog"
	"os"
	"time"

	_ "net/http/pprof"
	"repeatro/internal/config"
	"repeatro/internal/repositories"
	"repeatro/internal/security"
	"repeatro/internal/server"

	"github.com/lmittmann/tint"
)

var logger *slog.Logger

func main() {
	config := config.InitConfig("config")

	security := security.Security{ExpirationDelta: 600 * time.Minute}
	security.GetKyes()
	
	db := repositories.InitDatabase(config)
	
	repositories.InitGooseMigration(db)
	
	server := server.InitHTTPServer(config, db, security)
	
	logger = slog.New(tint.NewHandler(os.Stdout, nil))
	/*
		NOTE: stuff for pprof
		go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	*/
	// Create a new logger

	// Set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stdout, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	logger.Error("doint something")

	server.StartHttpServer()
}
