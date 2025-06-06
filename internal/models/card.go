package models

import (
	_"errors"
	"time"

	_"gorm.io/gorm"
)

type Card struct {
	CardId          string    `gorm:"primaryKey;size:36;not null" json:"card_id"`
	Word            string    `gorm:"type:varchar(100);not null;default:null" json:"word"`
	Translation     string    `gorm:"type:varchar(100);not null;default:null" json:"translation"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	ExpiresAt       time.Time `gorm:"autoCreateTime" json:"expires_at"`
	RepetitionNumber int   `gorm:"type:smallint;not null;default=0" json:"repetition_number"`
}

// func (c *Card) BeforeCreate(tx *gorm.DB) error {
// 	if c.Word == "" {
// 		return errors.New("email must not be empty")
// 	}
// 	return nil
// }