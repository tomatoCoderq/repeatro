package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"repeatro/internal/models"
	"repeatro/internal/schemes"
	"repeatro/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CardController struct {
	CardService services.CardServiceInterface
}

func CreateNewCardController(cardService *services.CardService) *CardController {
	return &CardController{CardService: cardService}
}

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
		fmt.Print("was heere")
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (cc CardController) ReadAllCardsToLearn(ctx *gin.Context) {
	response, err := cc.CardService.ReadAllCards()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (cc CardController) UpdateCard(ctx *gin.Context) {
	card_id := ctx.Param("id")
	// card_id := ctx.Request.Header.Get("id")

	cardId, err := uuid.Parse(card_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var cardUpdate schemes.UpdateCardScheme

	if err = json.Unmarshal(data, &cardUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	card, err := cc.CardService.UpdateCard(cardId, &cardUpdate)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("CO", card)
	ctx.JSON(http.StatusOK, card)
}

// Delete
func (cc CardController) DeleteCard(ctx *gin.Context) {
	card_id := ctx.Request.URL.Query().Get("card_id")

	cardId, err := uuid.Parse(card_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err = cc.CardService.DeleteCard(cardId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Status(http.StatusOK)
}

func (cc CardController) AddAnswers(ctx *gin.Context) {
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	var answers []schemes.AnswerScheme
	if err = json.Unmarshal(data, &answers); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	if err = cc.CardService.AddAnswers(answers); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	ctx.JSON(200, gin.H{"message": "added answers succesfully "})
}
