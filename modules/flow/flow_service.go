package flow

import (
	"moneh/models"
	"moneh/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Flow Interface
type FlowService interface {
	GetAllFlow(pagination utils.Pagination, userID uuid.UUID) ([]models.Flow, int, error)
	SoftDeleteFlowById(userID, flowID uuid.UUID) error
}

// Flow Struct
type flowService struct {
	flowRepo FlowRepository
}

// Flow Constructor
func NewFlowService(flowRepo FlowRepository) FlowService {
	return &flowService{
		flowRepo: flowRepo,
	}
}

func (r *flowService) GetAllFlow(pagination utils.Pagination, userID uuid.UUID) ([]models.Flow, int, error) {
	return r.flowRepo.FindAllFlow(pagination, userID)
}

func (s *flowService) SoftDeleteFlowById(userID, flowID uuid.UUID) error {
	// Repo : Find Flow By Id
	flow, err := s.flowRepo.FindFlowById(flowID)
	if err != nil {
		return err
	}
	if flow.DeletedAt != nil {
		return gorm.ErrRecordNotFound
	}

	// Soft Delete
	now := time.Now()
	flow.DeletedAt = &now

	// Repo : Update Flow By Id
	err = s.flowRepo.UpdateFlowById(flow, flowID)
	if err != nil {
		return err
	}

	return nil
}
