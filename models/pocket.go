package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	Pocket struct {
		ID            uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
		PocketName    string         `json:"pocket_name" gorm:"type:varchar(36);not null" binding:"required,max=36"`
		PocketDesc    *string        `json:"pocket_desc" gorm:"type:varchar(255);null" binding:"omitempty,max=255"`
		PocketType    string         `json:"pocket_type" gorm:"type:varchar(36);not null" binding:"required,max=36"`
		PocketAmmount int            `json:"pocket_ammount" gorm:"type:int;not null" binding:"required,min=0,max=9999999999"`
		PocketLimit   *int           `json:"pocket_limit" gorm:"type:int;null" binding:"omitempty,min=100000,max=99999999"`
		PocketTags    datatypes.JSON `json:"pocket_tags" gorm:"type:json;null"`
		CreatedAt     time.Time      `json:"created_at" gorm:"type:timestamp;not null"`
		UpdatedAt     *time.Time     `json:"updated_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      *User     `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"-"`
	}
)
