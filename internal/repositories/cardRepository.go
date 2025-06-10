package repositories

import (
	"fmt"
	"strconv"
	"time"

	"repeatro/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

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
	UpdateCard(id int) (models.Card, error)
	DeleteCard(id int) error
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

func (cr CardRepository) UpdateCard(id int) (models.Card, error) {
	return models.Card{}, nil
}

func (cr CardRepository) DeleteCard(id int) error {
	strId := strconv.Itoa(id)
	fmt.Println(strId)
	err := cr.db.Delete(&models.Card{}, "card_id = ?", strId).Error
	return err
}

// <---Mock functions--->

func (cm CardRepositoryMock) AddCard(card *models.Card) error {
	return nil
}

func (cm CardRepositoryMock) ReadAllCards() ([]models.Card, error) {
	return []models.Card{{CardId: uuid.New()}, {CardId: uuid.New()}}, nil
}

func (cm CardRepositoryMock) UpdateCard(id int) (models.Card, error) {
	return models.Card{CardId: uuid.New(), RepetitionNumber: 1}, nil
}

func (cm CardRepositoryMock) DeleteCard(id int) error {
	return nil
}
