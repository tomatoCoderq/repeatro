package server

import (
	"net/http"
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
	router := gin.Default()

	router.Handle(http.MethodPost, "cards/", cardController.AddCard)
	router.Handle(http.MethodGet, "cards/", cardController.ReadAllCardsToLearn)
	router.Handle(http.MethodPost, "cards/delete/", cardController.DeleteCard)

	return &HttpServer{
		config:         config,
		router:         router,
		db:             db,
		cardController: cardController,
	}
}

func (hs *HttpServer) StartHttpServer() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		panic("error during starting")
	}
}
