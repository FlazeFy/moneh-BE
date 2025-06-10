package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	History struct {
		ID           uuid.UUID  `json:"id" gorm:"type:varchar(36);primaryKey"`
		AdminID      *uuid.UUID `json:"admin_id" gorm:"type:varchar(36);null"`
		TechnicianID *uuid.UUID `json:"technician_id" gorm:"type:varchar(36);null"`
		UserID       *uuid.UUID `json:"user_id" gorm:"type:varchar(36);null"`
		TypeUser     string     `json:"type_user" gorm:"type:varchar(36);not null"`
		TypeHistory  string     `json:"type_history" gorm:"type:varchar(255);not null"`
		CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	}
	AllHistory struct {
		ID          uuid.UUID `json:"id" gorm:"type:varchar(36);primaryKey"`
		TypeUser    string    `json:"type_user" gorm:"type:varchar(36);not null"`
		TypeHistory string    `json:"type_history" gorm:"type:varchar(255);not null"`
		CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
		// FK - Admin
		AdminID *uuid.UUID `json:"admin_id"`
		Admin   *Admin     `json:"admins" gorm:"foreignKey:AdminID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		// FK - User
		UserID *uuid.UUID `json:"user_id"`
		User   *User      `json:"users" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
