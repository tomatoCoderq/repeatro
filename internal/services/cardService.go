package services

import (
	"fmt"
	"strconv"
	"time"
	"unicode"

	"repeatro/internal/models"
	"repeatro/internal/repositories"
)

func IsStringDigit(input string) bool {
	for _, i := range input {
		if !unicode.IsDigit(i) {
			return false
		}
	}
	return true
}

func SMTwoAlgo(currentDate time.Time, grade int, repetitionNumber int, easinessFactor float32, interval time.Duration) {
}

type CardService struct {
	cardRepository repositories.CardRepositoryInterface
}

type CardServiceMock struct{}

func CreateNewCardService(cardRepository *repositories.CardRepository) *CardService {
	return &CardService{cardRepository: cardRepository}
}

type CardServiceInterface interface {
	AddCard(card *models.Card) (*models.Card, error)
	ReadAllCards() ([]models.Card, error)
	UpdateCard(id string) (models.Card, error)
	DeleteCard(id string) error
}

func (cs CardService) AddCard(card *models.Card) (*models.Card, error) {
	// check all elements, write via rep
	if !IsStringDigit(card.CardId) {
		return nil, fmt.Errorf("CardId is not a digit")
	}

	card.ExpiresAt = time.Now().Add(10 * time.Second)

	err := cs.cardRepository.AddCard(card)
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (cm CardService) ReadAllCards() ([]models.Card, error) {
	cards, err := cm.cardRepository.ReadAllCards()
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (cm CardService) UpdateCard(id string) (models.Card, error) {
	if !IsStringDigit(id) {
		return models.Card{}, fmt.Errorf("CardId is not a digit")
	}

	idInt, _ := strconv.Atoi(id)

	card, err := cm.cardRepository.UpdateCard(idInt)
	if err != nil {
		return models.Card{}, nil
	}

	return card, nil
}

func (cm CardService) DeleteCard(id string) error {
	if !IsStringDigit(id) {
		return fmt.Errorf("CardId is not a digit")
	}

	idInt, _ := strconv.Atoi(id)

	err := cm.cardRepository.DeleteCard(idInt)
	if err != nil {
		return nil
	}

	return nil
}

func (cm CardServiceMock) AddCard(card *models.Card) (*models.Card, error) {
	return &models.Card{}, nil
}

func (cm CardServiceMock) ReadAllCards() ([]models.Card, error) {
	return []models.Card{{CardId: "idid"}}, nil
}

func (cm CardServiceMock) UpdateCard(id string) (models.Card, error) {
	return models.Card{CardId: "idid"}, nil
}

func (cm CardServiceMock) DeleteCard(id string) error {
	return nil
}
