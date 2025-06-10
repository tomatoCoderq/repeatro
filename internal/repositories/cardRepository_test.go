package repositories

import (
	"log"
	"strconv"
	"testing"
	"time"

	"repeatro/internal/models"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/google/uuid"

	// "gorm.io/gorm/logger"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	return gormDB, mock
}

func TestAddCard(t *testing.T) {
	db, mock := NewMockDB()
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	repo := CreateNewCardRepository(db)

	card := &models.Card{
		CardId:      uuid.New(),
		Word:        "hello",
		Translation: "hola",
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO .*").
		WithArgs(
			card.CardId,
			sqlmock.AnyArg(), // created_at
			sqlmock.AnyArg(), // expires_at
			0,                // repetition_number
			card.Word,
			card.Translation,
		).
		WillReturnRows(sqlmock.NewRows([]string{"word", "translation"}).AddRow(card.Word, card.Translation))
	mock.ExpectCommit()

	err := repo.AddCard(card)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestReadAllCards(t *testing.T) {
	db, mock := NewMockDB()
	repo := CardRepository{db: db}

	now := time.Now()

	rows := sqlmock.NewRows([]string{
		"card_id", "word", "translation", "created_at", "expires_at", "repetition_number",
	}).AddRow(
		"1", "hello", "hola", now, now.Add(-1*time.Hour), "1",
	).AddRow(
		"2", "bye", "adios", now, now.Add(-2*time.Hour), "2",
	)

	mock.ExpectQuery(`SELECT \* FROM "cards" WHERE expires_at < \$1`).
		WithArgs(sqlmock.AnyArg()).
		WillReturnRows(rows)

	cards, err := repo.ReadAllCards()

	assert.NoError(t, err)
	assert.Len(t, cards, 2)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteCard(t *testing.T) {
	db, mock := NewMockDB()
	repo := CardRepository{db: db}

	cardID := "123"
	id, _ := strconv.Atoi(cardID)
	// word := "test"
	// translation := "test translation"

	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM .* WHERE card_id = \$1`).
		WithArgs(cardID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.DeleteCard(id)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
