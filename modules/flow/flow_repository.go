package flow

import (
	"moneh/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Flow Interface
type FlowRepository interface {
	CreateFlow(flow *models.Flow, userID uuid.UUID) error

	// For Seeder
	DeleteAll() error
	FindOneRandomByUserID(userID uuid.UUID) (*models.Flow, error)
}

// Flow Struct
type flowRepository struct {
	db *gorm.DB
}

// Flow Constructor
func NewFlowRepository(db *gorm.DB) FlowRepository {
	return &flowRepository{db: db}
}

func (r *flowRepository) CreateFlow(flow *models.Flow, userID uuid.UUID) error {
	// Default
	flow.ID = uuid.New()
	flow.CreatedAt = time.Now()
	flow.CreatedBy = userID
	flow.UpdatedAt = nil
	flow.DeletedAt = nil

	// Query
	return r.db.Create(flow).Error
}

// For Seeder
func (r *flowRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Flow{}).Error
}

func (r *flowRepository) FindOneRandomByUserID(userID uuid.UUID) (*models.Flow, error) {
	var flow models.Flow

	err := r.db.Where("created_by", userID).Order("RANDOM()").Limit(1).First(&flow).Error
	if err != nil {
		return nil, err
	}

	return &flow, nil
}
