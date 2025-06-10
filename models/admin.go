package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Admin struct {
		ID              uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
		Username        string    `json:"username" gorm:"type:varchar(36);not null"`
		Password        string    `json:"password" gorm:"type:varchar(500);not null"`
		Email           string    `json:"email" gorm:"type:varchar(500);not null"`
		TelegramUserId  *string   `json:"telegram_user_id" gorm:"type:varchar(36);null"`
		TelegramIsValid bool      `json:"telegram_is_valid"`
		CreatedAt       time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	}
)

// For Generic Interface
func (a *Admin) GetID() uuid.UUID {
	return a.ID
}
func (a *Admin) GetPassword() string {
	return a.Password
}
