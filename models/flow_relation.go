package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	FlowRelation struct {
		ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		Ammount   int       `json:"ammount" gorm:"type:int;not null"`
		CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"-" binding:"-"`
		// FK - Flow
		FlowId uuid.UUID `json:"flow_id" gorm:"not null"`
		Flow   Flow      `json:"-" gorm:"foreignKey:FlowId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"-" binding:"-"`
		// FK - Pocket
		PocketId uuid.UUID `json:"pocket_id" gorm:"not null" binding:"required,max=36,min=36"`
		Pocket   Pocket    `json:"-" gorm:"foreignKey:PocketId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"-" binding:"-"`
	}
)
