package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	Pocket struct {
		ID            uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
		PocketName    string         `json:"pocket_name" gorm:"type:varchar(36);not null"`
		PocketDesc    *string        `json:"pocket_desc" gorm:"type:varchar(255);null"`
		PocketType    string         `json:"pocket_type" gorm:"type:varchar(36);not null"`
		PocketAmmount int            `json:"pocket_ammount" gorm:"type:int;not null"`
		PocketLimit   *int           `json:"pocket_limit" gorm:"type:int;null"`
		PocketTags    datatypes.JSON `json:"pocket_tags" gorm:"type:json;null"`
		CreatedAt     time.Time      `json:"created_at" gorm:"type:timestamp;not null"`
		UpdatedAt     *time.Time     `json:"updated_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
