package models

import (
	_ "errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Card struct {
	CardId          uuid.UUID    `gorm:"type:uuid;primaryKey;" json:"card_id"`
	Word            string    `gorm:"type:varchar(100);not null;default:null" json:"word"`
	Translation     string    `gorm:"type:varchar(100);not null;default:null" json:"translation"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	ExpiresAt       time.Time `gorm:"autoCreateTime" json:"expires_at"`
	RepetitionNumber int   `gorm:"type:smallint;default=0" json:"repetition_number"`
}

func (c *Card) BeforeCreate(tx *gorm.DB) error {
	c.CardId = uuid.New()
	c.ExpiresAt = time.Now().Add(10 * time.Second)
	return nil
}