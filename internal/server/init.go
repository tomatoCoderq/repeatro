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
	cardRepository := repositories.CreateNewCardRepository(db)
	cardService := services.CreateNewCardService(cardRepository)
	cardController := controllers.CreateNewCardController(cardService)

	deckRepository := repositories.CreateNewDeckRepository(db)
	deckService := services.CreateNewDeckService(deckRepository, cardRepository)
	deckController := controllers.CreateNewDeckController(deckService)

	userRepository := repositories.CreateNewUserRepository(db)
	userService := services.CreateNewUserService(userRepository, &security)
	userController := controllers.CreateNewUserController(userService, &security)

	security.UserRepository = userRepository

	// here routers
	router := gin.Default()
	router.RedirectFixedPath = false

	/* NOTE: Actually maybe i can not verify token each tome somehow and cache or smth*/

	secured := router.Group("/cards")
	securedDecks := router.Group("/decks")

	secured.Use(security.AuthMiddleware(), middlewares.ValidUserMiddleware())
	securedDecks.Use(security.AuthMiddleware())


	/* TODO: Need to check again all deck connected functions */
	securedDecks.Handle(http.MethodPost, "", deckController.AddDeck)
	securedDecks.Handle(http.MethodGet, "", deckController.ReadAllDecks)
	securedDecks.Handle(http.MethodGet, "/:id", deckController.ReadDeck)
	securedDecks.Handle(http.MethodDelete, "/:id", deckController.DeleteDeck)
	securedDecks.Handle(http.MethodPost, "/:deck_id/cards/:card_id", deckController.AddCardToDeck) // post one card
	securedDecks.Handle(http.MethodGet, "/:id/cards", deckController.ReadCardsFromDeck)

	secured.Handle(http.MethodPost, "", cardController.AddCard)
	// BUG: Some bug here with update_at and expires_at. Recheck and fix
	secured.Handle(http.MethodPost, "/answers", cardController.AddAnswers)
	secured.Handle(http.MethodPut, "/:id", cardController.UpdateCard)
	secured.Handle(http.MethodGet, "", cardController.ReadAllCardsToLearn)
	secured.Handle(http.MethodDelete, "/:id", cardController.DeleteCard)

	router.Handle(http.MethodPost, "register/", userController.Register)
	router.Handle(http.MethodPost, "login/", userController.Login)

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
