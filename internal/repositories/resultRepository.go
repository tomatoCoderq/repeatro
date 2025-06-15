package repositories

import (
	"fmt"
	"repeatro/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ResultRepository struct {
	db *gorm.DB
}

func CreateNewResultRepository(db *gorm.DB) *ResultRepository {
	return &ResultRepository{db: db}
}

type ResultRepositoryInterface interface {
	/* mean grade some period / cards learned by date / Add results / Delete results*/
	AddResult(result *models.Result) error
	DeleteResult(resultId uuid.UUID) error
	// same as next
	GetAllGradesForPeriod(dtStart time.Time, dtEnd time.Time, userId uuid.UUID) ([]int, error) 
	// Here i basically get all card for specific user over a period 
	GetLearnedCardsForPeriod(dtStart time.Time, dtEnd time.Time, userId uuid.UUID) ([]uuid.UUID, error) 
}

func (r *ResultRepository) AddResult(result *models.Result) error {
	return r.db.Create(result).Error
}

func (r *ResultRepository) DeleteResult(resultId uuid.UUID) error {
	return r.db.Delete(&models.Result{}, "id = ?", resultId).Error
}

func (r *ResultRepository) GetAllGradesForPeriod(dtStart, dtEnd time.Time, userId uuid.UUID) ([]int, error) {
	var grades []int
	err := r.db.
		Model(&models.Result{}).
		Where("created_at BETWEEN ? AND ?", dtStart, dtEnd).
		Pluck("grade", &grades).Error
	fmt.Println("so", grades, dtStart, dtEnd)
	return grades, err
}

func (r *ResultRepository) GetLearnedCardsForPeriod(dtStart, dtEnd time.Time, userId uuid.UUID) ([]uuid.UUID, error) {
	var cardIDs []uuid.UUID
	err := r.db.
		Model(&models.Result{}).
		Select("DISTINCT card_id").
		Where("user_id = ? AND created_at BETWEEN ? AND ?", userId, dtStart, dtEnd).
		Pluck("card_id", &cardIDs).Error
	return cardIDs, err
}



