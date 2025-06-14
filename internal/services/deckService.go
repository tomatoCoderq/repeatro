package services

import (
	"errors"
	"repeatro/internal/models"
	"repeatro/internal/repositories"

	"github.com/google/uuid"
)

var ErrUnauthorized = errors.New("you do not own this deck")

type DeckService struct {
	deckRepository repositories.DeckRepositoryInterface
	cardRepository repositories.CardRepositoryInterface
}

func CreateNewDeckService(deckRepository *repositories.DeckRepository, cardRepository *repositories.CardRepository) *DeckService {
	return &DeckService{
		deckRepository: deckRepository,
		cardRepository: cardRepository,
	}
}

type DeckServiceInterface interface {
	AddCard(deck *models.Deck, userId uuid.UUID) (*models.Deck, error)
	ReadAllDecksOfUser(userId uuid.UUID) ([]models.Deck, error)
	ReadAllCardsFromDeck(deckId uuid.UUID, userId uuid.UUID) ([]models.Card, error)
	ReadDeck(deckId uuid.UUID, userId uuid.UUID) (*models.Deck, error)
	DeleteDeck(deckId uuid.UUID, userId uuid.UUID) error
	AddCardToDeck(cardId uuid.UUID, deckId uuid.UUID, userId uuid.UUID) error
}

func (ds *DeckService) AddCard(deck *models.Deck, userId uuid.UUID) (*models.Deck, error) {
	deck.CreatedBy = userId
	err := ds.deckRepository.AddDeck(deck)
	if err != nil {
		return nil, err
	}
	return deck, nil
}

func (ds *DeckService) ReadAllDecksOfUser(userId uuid.UUID) ([]models.Deck, error) {
	return ds.deckRepository.ReadAllDecksOfUser(userId)
}

func (ds *DeckService) ReadDeck(deckId uuid.UUID, userId uuid.UUID) (*models.Deck, error) {
	deck, err := ds.deckRepository.ReadDeck(deckId)
	if err != nil {
		return nil, err
	}

	if deck.CreatedBy != userId {
		return nil, ErrUnauthorized
	}
	return deck, nil
}

func (ds *DeckService) ReadAllCardsFromDeck(deckId uuid.UUID, userId uuid.UUID) ([]models.Card, error) {
	// deck, err := ds.ReadDeck(deckId, userId)
	// if err != nil {
	// 	return nil, err
	// }
	// return deck.Cards, nil
	cards, err := ds.deckRepository.FindAllCardsInDeck(deckId)
	if err != nil {
		return nil, err
	}
	return cards, nil
	
}
  
func (ds *DeckService) DeleteDeck(deckId uuid.UUID, userId uuid.UUID) error {
	deck, err := ds.deckRepository.ReadDeck(deckId)
	if err != nil {
		return err
	}
	if deck.CreatedBy != userId {
		return ErrUnauthorized
	}
	return ds.deckRepository.DeleteDeck(deckId)
}

func (ds *DeckService) AddCardToDeck(cardId uuid.UUID, deckId uuid.UUID, userId uuid.UUID) error {
	deck, err := ds.deckRepository.ReadDeck(deckId)
	if err != nil {
		return err
	}
	if deck.CreatedBy != userId {
		return ErrUnauthorized
	}
	return ds.deckRepository.AddCardToDeck(cardId, deckId)
}
