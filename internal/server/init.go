package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"repeatro/internal/controllers"
	"repeatro/internal/middlewares"
	"repeatro/internal/repositories"
	"repeatro/internal/security"

	// "github.com/swaggo/files"       // swagger embed files
	// "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"repeatro/internal/services"
)

type HttpServer struct {
	config         *viper.Viper
	router         *gin.Engine
	db             *gorm.DB
	cardController *controllers.CardController
}

func InitHTTPServer(config *viper.Viper, db *gorm.DB, security security.Security) *HttpServer {
	// TODO: Add logger as middleware

	resultRepository := repositories.CreateNewResultRepository(db)
	cardRepository := repositories.CreateNewCardRepository(db)
	deckRepository := repositories.CreateNewDeckRepository(db)
	userRepository := repositories.CreateNewUserRepository(db)

	cardService := services.CreateNewCardService(cardRepository, resultRepository)
	resultsService := services.CreateNewResultService(resultRepository, cardRepository)
	deckService := services.CreateNewDeckService(deckRepository, cardRepository)
	userService := services.CreateNewUserService(userRepository, &security)

	resultContoller := controllers.CreateNewResultController(resultsService)
	cardController := controllers.CreateNewCardController(cardService)
	deckController := controllers.CreateNewDeckController(deckService)
	userController := controllers.CreateNewUserController(userService, &security)

	security.UserRepository = userRepository

	// here routers
	router := gin.Default()
	router.RedirectFixedPath = false

	/* NOTE: Actually maybe i can not verify token each tome somehow and cache or smth*/
	secured := router.Group("")
	secured.Use(gin.Recovery(), security.AuthMiddleware(), middlewares.ValidUserMiddleware())

	cards := secured.Group("/cards")
	decks := secured.Group("/decks")
	stats := secured.Group("/stats")

	router.Handle(http.MethodPost, "register/", userController.Register)
	router.Handle(http.MethodPost, "login/", userController.Login)

	cards.Handle(http.MethodPost, "", cardController.AddCard)
	// BUG: Some bug here with update_at and expires_at. Recheck and fix
	cards.Handle(http.MethodGet, "", cardController.ReadAllCardsToLearn)
	cards.Handle(http.MethodPut, "/:id", cardController.UpdateCard)
	cards.Handle(http.MethodDelete, "/:id", cardController.DeleteCard)
	cards.Handle(http.MethodPost, "/answers", cardController.AddAnswers)

	decks.Handle(http.MethodPost, "", deckController.AddDeck)
	decks.Handle(http.MethodGet, "", deckController.ReadAllDecks)
	decks.Handle(http.MethodGet, "/:id", deckController.ReadDeck)
	decks.Handle(http.MethodDelete, "/:id", deckController.DeleteDeck)
	decks.Handle(http.MethodPost, "/:deck_id/cards/:card_id", deckController.AddCardToDeck) // post one card
	decks.Handle(http.MethodGet, "/:id/cards", deckController.ReadCardsFromDeck)

	stats.Handle(http.MethodGet, "", resultContoller.GetStats)

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
