package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	Flow struct {
		ID            uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
		FlowType      string         `json:"flow_type" gorm:"type:varchar(36);not null" binding:"required,max=36"`
		FlowCategory  string         `json:"flow_category" gorm:"type:varchar(36);not null" binding:"required,max=36"`
		FlowName      string         `json:"flow_name" gorm:"type:varchar(144);not null" binding:"required,max=144"`
		FlowDesc      *string        `json:"flow_desc" gorm:"type:varchar(255);null" binding:"omitempty,max=255"`
		FlowTag       datatypes.JSON `json:"flow_tags" gorm:"type:json;null"`
		IsSplitBill   bool           `json:"is_split_bill" gorm:"type:boolean;not null" binding:"required"`
		IsMultiPocket bool           `json:"is_multi_pocket" gorm:"type:boolean;not null" binding:"required"`
		CreatedAt     time.Time      `json:"created_at" gorm:"type:timestamp;not null"`
		UpdatedAt     *time.Time     `json:"updated_at" gorm:"type:timestamp;null"`
		DeletedAt     *time.Time     `json:"deleted_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" validate:"-"`

		FlowRelations []FlowRelation `json:"flow_relations" gorm:"foreignKey:flow_id;references:id"`
	}
	FlowPocketTags struct {
		TagName string `json:"tag_name"`
		TagSlug string `json:"tag_slug"`
	}
)
