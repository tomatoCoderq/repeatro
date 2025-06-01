package repositories

import (
	"strconv"

	"repeatro/models"

	"gorm.io/gorm"
)

// TODO: write repository part, choose bd(postgresql?)
type CardRepository struct {
	db *gorm.DB
}

func CreateNewCardRepository(db *gorm.DB) *CardRepository{
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
	result := cr.db.Create(card)
	return result.Error
}

func (cm CardRepository) ReadAllCards() ([]models.Card, error) {
	return []models.Card{}, nil
}

func (cm CardRepository) UpdateCard(id int) (models.Card, error) {
	return models.Card{}, nil
}

func (cm CardRepository) DeleteCard(id int) error {
	return nil
}

func (cm CardRepositoryMock) AddCard(card *models.Card) error {
	return nil
}

func (cm CardRepositoryMock) ReadAllCards() ([]models.Card, error) {
	return []models.Card{{CardId: "1"}, {CardId: "2"}}, nil
}

func (cm CardRepositoryMock) UpdateCard(id int) (models.Card, error) {
	return models.Card{CardId: strconv.Itoa(id), RepetitionNumber: "1"}, nil
}

func (cm CardRepositoryMock) DeleteCard(id int) error {
	return nil
}
