package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"repeatro/internal/controllers"
	"repeatro/internal/repositories"
	"repeatro/internal/services"
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
	router.Handle(http.MethodDelete, "cards/", cardController.DeleteCard)

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
