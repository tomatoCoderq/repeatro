package models

import "time"

type Card struct {
	CardId          string    `gorm:"primaryKey;size:36;not null" json:"card_id"`
	Word            string    `gorm:"type:varchar(100);not null" json:"word"`
	Translation     string    `gorm:"type:varchar(100);not null" json:"translation"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	ExpiresAt       time.Time `gorm:"not null" json:"expires_at"`
	RepetitionNumber string   `gorm:"type:varchar(10);not null" json:"repetition_number"`
}