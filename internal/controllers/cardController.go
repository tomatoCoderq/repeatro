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
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func getUserIdFromContext(ctx *gin.Context) (uuid.UUID, error) {
	userClaims, exists := ctx.Get("userClaims")
	if !exists {
		return uuid.UUID{}, fmt.Errorf("user claims do not exist")
	}

	claimsMap, ok := userClaims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("cannot convert claims to map")
	}

	userIdString, ok := claimsMap["user_id"].(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("cannot get user_id from claims map")
	}

	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return uuid.UUID{}, err
	}
	return userId, nil
}

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

	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var card models.Card
	if err = json.Unmarshal(body, &card); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	card.CreatedBy = userId

	response, err := cc.CardService.AddCard(&card)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (cc CardController) ReadAllCardsToLearn(ctx *gin.Context) {
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response, err := cc.CardService.ReadAllCards(userId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

func (cc CardController) UpdateCard(ctx *gin.Context) {
	card_id := ctx.Param("id")
	cardId, err := uuid.Parse(card_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
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

	card, err := cc.CardService.UpdateCard(cardId, &cardUpdate, userId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println("CO", card)
	ctx.JSON(http.StatusOK, card)
}

// Delete
func (cc CardController) DeleteCard(ctx *gin.Context) {
	card_id := ctx.Param("id")
	cardId, err := uuid.Parse(card_id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = cc.CardService.DeleteCard(cardId, userId)
	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.Status(http.StatusOK)
}

func (cc CardController) AddAnswers(ctx *gin.Context) {
	data, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var answers []schemes.AnswerScheme
	if err = json.Unmarshal(data, &answers); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err = cc.CardService.AddAnswers(userId, answers); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "added answers succesfully "})
}
