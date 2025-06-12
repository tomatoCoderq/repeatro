package repositories

import (
	"fmt"
	"time"

	"repeatro/internal/models"
	"repeatro/internal/schemes"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func updateCardFields(cardInitial *models.Card, card *schemes.UpdateCardScheme) {
	if card.Word != "" {
		cardInitial.Word = card.Word
	}
	if card.Translation != "" {
		cardInitial.Translation = card.Translation
	}
	if card.Easiness != 0 {
		cardInitial.Easiness = card.Easiness
	}
	if !card.UpdatedAt.IsZero() {
		cardInitial.UpdatedAt = card.UpdatedAt
	}
	if card.Interval != 0 {
		cardInitial.Interval = card.Interval
	}
	if !card.ExpiresAt.IsZero() {
		cardInitial.ExpiresAt = card.ExpiresAt
	}
	if card.RepetitionNumber != 0 {
		cardInitial.RepetitionNumber = card.RepetitionNumber
	}
	fmt.Println("CHA:", cardInitial)
}

// TODO: write repository part, choose bd(postgresql?)
type CardRepository struct {
	db *gorm.DB
}

func CreateNewCardRepository(db *gorm.DB) *CardRepository {
	return &CardRepository{db: db}
}

type CardRepositoryMock struct{}

type CardRepositoryInterface interface {
	AddCard(card *models.Card) error
	ReadAllCards() ([]models.Card, error)
	ReadCard(cardId uuid.UUID) (*models.Card, error)
	PureUpdate(card *models.Card) error
	UpdateCard(card *models.Card, cardUpdate *schemes.UpdateCardScheme) (*models.Card, error)
	DeleteCard(cardId uuid.UUID) error
	// AddAnswers(answers []schemes.Answer) error
}

func (cr CardRepository) AddCard(card *models.Card) error {
	return cr.db.Create(card).Error	
}

func (cr CardRepository) ReadAllCards() ([]models.Card, error) {
	var cards []models.Card
	err := cr.db.Where("expires_at < ?", time.Now()).Find(&cards).Error
	if err != nil {
		return nil, err
	}
	return cards, err
}

func (cr CardRepository) ReadCard(cardId uuid.UUID) (*models.Card, error) {
	var card models.Card
	err := cr.db.Where("card_id = ?", cardId).Find(&card).Error
	return &card, err
}

func (cr CardRepository) UpdateCard(card *models.Card, cardUpdate *schemes.UpdateCardScheme) (*models.Card, error) {
	updateCardFields(card, cardUpdate)	
	return card, cr.db.Updates(card).Error
}

func (cr CardRepository) PureUpdate(card *models.Card) error {
	fmt.Println(card)
	return cr.db.Updates(card).Error
}

func (cr CardRepository) DeleteCard(cardId uuid.UUID) error {
	err := cr.db.Delete(&models.Card{}, "card_id = ?", cardId).Error
	return err
}

// <---Mock functions--->

func (cm CardRepositoryMock) AddCard(card *models.Card) error {
	return nil
}

func (cm CardRepositoryMock) ReadAllCards() ([]models.Card, error) {
	return []models.Card{{CardId: uuid.New()}, {CardId: uuid.New()}}, nil
}

func (cm CardRepositoryMock) ReadCard(cardId uuid.UUID) (*models.Card, error) {
	return nil, nil
}

func (cm CardRepositoryMock) UpdateCard(card *models.Card) (*models.Card, error) {
	return &models.Card{CardId: uuid.New(), RepetitionNumber: 1}, nil
}

func (cm CardRepositoryMock) DeleteCard(cardId uuid.UUID) error {
	return nil
}
