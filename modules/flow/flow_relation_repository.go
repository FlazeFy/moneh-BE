package flow

import (
	"moneh/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Flow Relation Interface
type FlowRelationRepository interface {
	CreateFlowRelation(flowRelation *models.FlowRelation, userID uuid.UUID) error

	// For Seeder
	DeleteAll() error
}

// Flow Relation Struct
type flowRelationRepository struct {
	db *gorm.DB
}

// Flow Relation Constructor
func NewFlowRelationRepository(db *gorm.DB) FlowRelationRepository {
	return &flowRelationRepository{db: db}
}

func (r *flowRelationRepository) CreateFlowRelation(flowRelation *models.FlowRelation, userID uuid.UUID) error {
	// Default
	flowRelation.ID = uuid.New()
	flowRelation.CreatedAt = time.Now()
	flowRelation.CreatedBy = userID

	// Query
	return r.db.Create(flowRelation).Error
}

// For Seeder
func (r *flowRelationRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.FlowRelation{}).Error
}
