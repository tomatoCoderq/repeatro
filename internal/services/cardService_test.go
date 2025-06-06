package services

import (
	"testing"
	"time"

	"repeatro/internal/models"
	"repeatro/internal/repositories"

	"github.com/stretchr/testify/assert"
)

func TestAddCard(t *testing.T) {
	cardService := CardService{cardRepository: repositories.CardRepositoryMock{}}
	card := &models.Card{}
	cardNew, err := cardService.AddCard(card)

	assert.Equal(t, nil, err)

	cardT := &models.Card{CreatedAt: time.Date(2022, time.April, 1, 12, 12, 12, 12, time.Now().Location()), RepetitionNumber: 0, ExpiresAt: card.CreatedAt.Add(10 * time.Second)}

	assert.Equal(t, cardNew.RepetitionNumber, cardT.RepetitionNumber)
}

func TestReadAllCards(t *testing.T) {
	// cardService := CardService{cardRepository: repositories.CardRepositoryMock{}}
	// cards, err := cardService.ReadAllCards()

	// assert.Equal(t, nil, err)

	// assert.Equal(t, []models.Card{{CardId: "1"}, {CardId: "2"}}, cards)
}

func TestUpdateCard(t *testing.T) {
	cardService := CardService{cardRepository: repositories.CardRepositoryMock{}}
	cardNew, err := cardService.UpdateCard("0")

	assert.Equal(t, nil, err)

	assert.Equal(t, models.Card{CardId: "0", RepetitionNumber: 1}, cardNew)
}

func TestDeleteCard(t *testing.T) {
	cardService := CardService{cardRepository: repositories.CardRepositoryMock{}}
	err := cardService.DeleteCard("0")

	assert.Equal(t, nil, err)
}
