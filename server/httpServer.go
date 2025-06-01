package server

import (
	"repeatro/controllers"
	"repeatro/repositories"
	"repeatro/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type HttpServer struct {
	config         *viper.Viper
	router         *gin.Engine
	db             *gorm.DB
	cardController *controllers.CardController
}

func InitHTTPServer(config *viper.Viper, db *gorm.DB) *HttpServer {
	cardRepository := repositories.CreateNewCardRepository(db)

	cardService := services.CreateNewCardService(cardRepository)

	cardController := controllers.CreateNewCardController(cardService)

	// here routers

	return &HttpServer{
		config:         config,
		router:         gin.Default(),
		db:             db,
		cardController: cardController,
	}
}

func (hs *HttpServer) StartHttpServer() {
	err := hs.router.Run(hs.config.GetString("http.server_addres"))
	if err != nil {
		panic("error during starting")
	}
}
