package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	// "repeatro/models"
	"repeatro/models"
	"repeatro/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func InitMock() CardController {
	cardServiceMock := services.CardServiceMock{}
	cardController := CardController{CardService: cardServiceMock}
	return cardController
}

// Insert
func TestAddCard(t *testing.T) {
	// cardService := services.CardService{}
	cardController := InitMock()

	cardJSON := `{"card_id": "idid", "word": "Bonjour"}`
	req := httptest.NewRequest(http.MethodPost, "/cards", bytes.NewBufferString(cardJSON))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	cardController.AddCard(c)
	assert.Equal(t, 200, w.Code)
}

// Read
func TestReadAllCardsToLearn(t *testing.T) {
	cardController := InitMock()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	cardController.ReadAllCardsToLearn(c)
	data := w.Body.Bytes()
	var cards []models.Card
	json.Unmarshal(data, &cards)
	assert.Equal(t, []models.Card{{CardId: "idid"}}, cards)
}


// // Update
func TestUpdateCard(t *testing.T) {
	cardController := InitMock()

	cardJSON := `{"card_id": "idid", "word": "BonjourYopta"}`
	req := httptest.NewRequest(http.MethodPost, "/cards/4", bytes.NewBufferString(cardJSON))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	cardController.UpdateCard(c)
	assert.Equal(t, 200, w.Code)

	data := w.Body.Bytes()
	var card models.Card
	json.Unmarshal(data, &card)

	assert.Equal(t, models.Card{CardId: "idid"}, card)
}

// // Delete
func TestDeleteCard(t *testing.T) {
	cardController := InitMock()

	cardJSON := `{"card_id": "idid", "word": "BonjourYopta"}`
	req := httptest.NewRequest(http.MethodDelete, "/cards/4", bytes.NewBufferString(cardJSON))

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	cardController.DeleteCard(c)
	assert.Equal(t, 200, w.Code)
}
