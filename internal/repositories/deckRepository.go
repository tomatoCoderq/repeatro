package repositories

import (
	"repeatro/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeckRepository struct {
	db *gorm.DB
}

func CreateNewDeckRepository(db *gorm.DB) *DeckRepository {
	return &DeckRepository{db: db}
}

type DeckRepositoryInterface interface {
	AddDeck(deck *models.Deck) error
	ReadAllDecksOfUser(userId uuid.UUID) ([]models.Deck, error)
	ReadAllDecks() ([]models.Deck, error)
	ReadDeck(deckId uuid.UUID) (*models.Deck, error)
	DeleteDeck(deckId uuid.UUID) error
	AddCardToDeck(cardId uuid.UUID, deckId uuid.UUID) error
}

func (r *DeckRepository) AddDeck(deck *models.Deck) error {
	return r.db.Create(deck).Error
}

func (r *DeckRepository) ReadAllDecksOfUser(userId uuid.UUID) ([]models.Deck, error) {
	var decks []models.Deck
	err := r.db.Where("created_by = ?", userId).Preload("Cards").Find(&decks).Error
	return decks, err
}

func (r *DeckRepository) ReadAllDecks() ([]models.Deck, error) {
	var decks []models.Deck
	err := r.db.Preload("Cards").Find(&decks).Error
	return decks, err
}

func (r *DeckRepository) ReadDeck(deckId uuid.UUID) (*models.Deck, error) {
	var deck models.Deck
	err := r.db.Where("id = ?", deckId).Preload("Cards").First(&deck).Error
	if err != nil {
		return nil, err
	}
	return &deck, nil
}

func (r *DeckRepository) DeleteDeck(deckId uuid.UUID) error {
	return r.db.Delete(&models.Deck{}, "id = ?", deckId).Error
}

func (r *DeckRepository) AddCardToDeck(cardId uuid.UUID, deckId uuid.UUID) error {
	return r.db.Model(&models.Card{}).Where("id = ?", cardId).Update("deck_id", deckId).Error
}
