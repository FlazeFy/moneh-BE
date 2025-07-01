package flow

import (
	"moneh/models"
	"moneh/utils"

	"github.com/google/uuid"
)

// Flow Interface
type FlowService interface {
	GetAllFlow(pagination utils.Pagination, userID uuid.UUID) ([]models.Flow, int, error)
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
