package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	Flow struct {
		ID            uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
		FlowType      string         `json:"flow_type" gorm:"type:varchar(36);not null"`
		FlowCategory  string         `json:"flow_category" gorm:"type:varchar(36);not null"`
		FlowName      string         `json:"flow_name" gorm:"type:varchar(144);not null"`
		FlowDesc      *string        `json:"flow_desc" gorm:"type:varchar(255);null"`
		FlowTag       datatypes.JSON `json:"flow_tags" gorm:"type:json;null"`
		IsSplitBill   bool           `json:"is_split_bill" gorm:"type:boolean;not null"`
		IsMultiPocket bool           `json:"is_multi_pocket" gorm:"type:boolean;not null"`
		CreatedAt     time.Time      `json:"created_at" gorm:"type:timestamp;not null"`
		UpdatedAt     *time.Time     `json:"updated_at" gorm:"type:timestamp;null"`
		DeletedAt     *time.Time     `json:"deleted_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
	FlowPocketTags struct {
		TagName string `json:"tag_name"`
		TagSlug string `json:"tag_slug"`
	}
)
