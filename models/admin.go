package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Admin struct {
		ID              uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
		Username        string    `json:"username" gorm:"type:varchar(36);not null" binding:"required,min=6,max=36"`
		Password        string    `json:"password" gorm:"type:varchar(500);not null" binding:"required,min=6,max=36"`
		Email           string    `json:"email" gorm:"type:varchar(500);not null" binding:"required,email,max=500"`
		TelegramUserId  *string   `json:"telegram_user_id" gorm:"type:varchar(36);null" binding:"omitempty,max=36"`
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
