package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"repeatro/models"
	"repeatro/services"

	"github.com/gin-gonic/gin"
)

type CardController struct {
	CardService services.CardServiceInterface
}

func CreateNewCardController(cardService *services.CardService) *CardController {
	return &CardController{CardService: cardService}
}

// Insert
func (cc CardController) AddCard(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var card models.Card
	if err = json.Unmarshal(body, &card); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, err := cc.CardService.AddCard(&card)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Read
func (cc CardController) ReadAllCardsToLearn(ctx *gin.Context) {
	response, err := cc.CardService.ReadAllCards()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Update
func (cc CardController) UpdateCard(ctx *gin.Context) {
	id := ctx.Request.Header.Get("id")

	card, err := cc.CardService.UpdateCard(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, card)
}

// Delete
func (cc CardController) DeleteCard(ctx *gin.Context) {
	id := ctx.Request.Header.Get("id")

	err := cc.CardService.DeleteCard(id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}
