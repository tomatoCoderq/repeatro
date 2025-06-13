package controllers

import (
	"net/http"

	"repeatro/internal/models"

	"repeatro/internal/services"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"repeatro/internal/tools"
)

type DeckController struct {
	DeckService services.DeckServiceInterface
}

func CreateNewDeckController(deckService *services.DeckService) *DeckController {
	return &DeckController{DeckService: deckService}
}

func (dc DeckController) AddDeck(ctx *gin.Context) {
	var deck models.Deck

	if err := ctx.ShouldBindJSON(&deck); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := tools.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deck.CreatedBy = userId
	createdDeck, err := dc.DeckService.AddCard(&deck, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, createdDeck)
}

func (dc DeckController) ReadAllDecks(ctx *gin.Context) {
	userId, err := tools.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	decks, err := dc.DeckService.ReadAllDecksOfUser(userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, decks)
}

func (dc DeckController) ReadDeck(ctx *gin.Context) {
	deckId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid deck ID"})
		return
	}

	userId, err := tools.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	deck, err := dc.DeckService.ReadDeck(deckId, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, deck)
}

// Delete
func (dc DeckController) DeleteDeck(ctx *gin.Context) {
	deckId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid deck ID"})
		return
	}

	userId, err := tools.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := dc.DeckService.DeleteDeck(deckId, userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (dc DeckController) AddCardToDeck(ctx *gin.Context) {
	cardId, err := uuid.Parse(ctx.Param("card_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid card ID"})
		return
	}

	deckId, err := uuid.Parse(ctx.Param("deck_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid deck ID"})
		return
	}

	userId, err := tools.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := dc.DeckService.AddCardToDeck(cardId, deckId, userId); err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "card added to deck"})
}
