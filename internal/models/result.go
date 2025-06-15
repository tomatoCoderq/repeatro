package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Result struct {
	ResultId uuid.UUID  `gorm:"type:uuid;primaryKey" json:"result_id"`
	UserId    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	User      User      `gorm:"foreignKey:UserId;references:ResultId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`

	CardId    uuid.UUID `gorm:"type:uuid;not null" json:"card_id"`
	Card      Card      `gorm:"foreignKey:CardId;references:ResultId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	Grade     int       `gorm:"not null" json:"grade"`
}

func (r *Result) BeforeCreate(tx *gorm.DB) error {
	r.ResultId = uuid.New()
	return nil
}


