package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"repeatro/internal/models"
	"repeatro/internal/repositories"
)

func getMean(grades []int) int {
	var sum int
	for _, grade := range grades {
		sum += grade
	}

	return sum / len(grades)
}

type ResultService struct {
	resultRepository repositories.ResultRepositoryInterface
	cardRepository   repositories.CardRepositoryInterface
}

func CreateNewResultService(resultRepository *repositories.ResultRepository, cardRepository *repositories.CardRepository) *ResultService {
	return &ResultService{
		resultRepository: resultRepository,
		cardRepository:   cardRepository,
	}
}

type ResultServiceInterface interface {
	// same as next
	GetMeanGradeOfPeriod(dtStart time.Time, dtEnd time.Time, userId uuid.UUID) (int, error)
	// Here i basically get all card for specific user over a period
	GetLearnedCardsForPeriod(dtStart time.Time, dtEnd time.Time, userId uuid.UUID) ([]*models.Card, error)
}

func (rs *ResultService) GetMeanGradeOfPeriod(dtStart time.Time, dtEnd time.Time, userId uuid.UUID) (int, error) {
	grades, err := rs.resultRepository.GetAllGradesForPeriod(dtStart, dtEnd, userId)
	if err != nil {
		return 0, err
	}

	fmt.Println("graaades", grades)

	if len(grades) == 0 {
		return 0, fmt.Errorf("grades over this period are not found")
	}

	return getMean(grades), nil
}

func (rs *ResultService) GetLearnedCardsForPeriod(dtStart time.Time, dtEnd time.Time, userId uuid.UUID) ([]*models.Card, error) {
	cardIds, err := rs.resultRepository.GetLearnedCardsForPeriod(dtStart, dtEnd, userId)
	if err != nil {
		return nil, err
	}

	if len(cardIds) == 0 {
		return nil, fmt.Errorf("learned cards over this period are not found")
	}

	cards := make([]*models.Card, 1)
	for _, cardId := range cardIds {
		card, err := rs.cardRepository.ReadCard(cardId)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil
}
