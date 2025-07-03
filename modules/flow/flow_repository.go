package flow

import (
	"moneh/models"
	"moneh/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Flow Interface
type FlowRepository interface {
	CreateFlow(flow *models.Flow, userID uuid.UUID) (*models.Flow, error)
	FindAllFlow(pagination utils.Pagination, userID uuid.UUID) ([]models.Flow, int, error)
	FindFlowById(ID uuid.UUID) (*models.Flow, error)
	UpdateFlowById(flow *models.Flow, ID uuid.UUID) error

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

func (r *flowRepository) FindFlowById(ID uuid.UUID) (*models.Flow, error) {
	// Model
	var flow models.Flow

	// Query
	result := r.db.Unscoped().Where("id = ?", ID).First(&flow)

	// Response
	if result.Error != nil {
		return nil, result.Error
	}

	return &flow, nil
}

func (r *flowRepository) UpdateFlowById(flow *models.Flow, ID uuid.UUID) error {
	now := time.Now()
	flow.ID = ID
	flow.UpdatedAt = &now

	if err := r.db.Save(flow).Error; err != nil {
		return err
	}

	return nil
}

func (r *flowRepository) FindAllFlow(pagination utils.Pagination, userID uuid.UUID) ([]models.Flow, int, error) {
	// Model
	var total int
	var flows []models.Flow

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit

	// Query
	err := r.db.Where("created_by = ? AND deleted_at IS NULL", userID).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset).
		Preload("FlowRelations").
		Find(&flows).Error

	if err != nil {
		return nil, 0, err
	}

	total = len(flows)
	return flows, total, nil
}

func (r *flowRepository) CreateFlow(flow *models.Flow, userID uuid.UUID) (*models.Flow, error) {
	flowRelationTemp := flow.FlowRelations

	// Default
	flow.ID = uuid.New()
	flow.CreatedAt = time.Now()
	flow.CreatedBy = userID
	flow.UpdatedAt = nil
	flow.DeletedAt = nil
	flow.FlowRelations = nil

	// Query
	if err := r.db.Create(flow).Error; err != nil {
		return nil, err
	}

	flow.FlowRelations = flowRelationTemp

	return flow, nil
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
