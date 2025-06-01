package models

import "time"

type Card struct {
	CardId string `json:"card_id"`
	Word string `json:"word"`
	Translation string `json:"translation"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
	RepetitionNumber string `json:"repetition_number"`
}