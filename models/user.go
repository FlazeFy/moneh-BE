package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	User struct {
		ID              uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
		Username        string    `json:"username" gorm:"type:varchar(36);not null" binding:"required,min=6,max=36"`
		Password        string    `json:"password" gorm:"type:varchar(500);not null" binding:"required,min=6,max=36"`
		Email           string    `json:"email" gorm:"type:varchar(500);not null" binding:"required,email,max=500"`
		Currency        string    `json:"currency" gorm:"type:varchar(3);not null" binding:"required,min=3,max=3"`
		TelegramUserId  *string   `json:"telegram_user_id" gorm:"type:varchar(36);null" binding:"omitempty,max=36"`
		TelegramIsValid bool      `json:"telegram_is_valid"`
		CreatedAt       time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	}
	UserAuth struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	// All Role
	Account interface {
		GetID() uuid.UUID
		GetPassword() string
	}
	MyProfile struct {
		Username        string    `json:"username" gorm:"type:varchar(36);not null"`
		Email           string    `json:"email" gorm:"type:varchar(500);not null"`
		TelegramUserId  *string   `json:"telegram_user_id" gorm:"type:varchar(36);null"`
		TelegramIsValid bool      `json:"telegram_is_valid"`
		Currency        string    `json:"currency" gorm:"type:varchar(3);not null"`
		CreatedAt       time.Time `json:"created_at" gorm:"type:timestamp;not null"`
	}
)

// For Generic Interface
func (a *User) GetID() uuid.UUID {
	return a.ID
}
func (a *User) GetPassword() string {
	return a.Password
}
