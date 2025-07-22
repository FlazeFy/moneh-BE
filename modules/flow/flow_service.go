package flow

import (
	"fmt"
	"moneh/models"
	"moneh/modules/history"
	"moneh/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Flow Interface
type FlowService interface {
	GetAllFlow(pagination utils.Pagination, userID uuid.UUID) ([]models.Flow, int, error)
	CreateFlow(flow *models.Flow, userID uuid.UUID) (*models.Flow, error)
	SoftDeleteFlowById(userID, flowID uuid.UUID) error
	HardDeleteFlowById(userID, flowID uuid.UUID) error
}

// Flow Struct
type flowService struct {
	flowRepo         FlowRepository
	historyRepo      history.HistoryRepository
	flowRelationRepo FlowRelationRepository
}

// Flow Constructor
func NewFlowService(flowRepo FlowRepository, historyRepo history.HistoryRepository, flowRelationRepo FlowRelationRepository) FlowService {
	return &flowService{
		flowRepo:         flowRepo,
		historyRepo:      historyRepo,
		flowRelationRepo: flowRelationRepo,
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

func (s *flowService) HardDeleteFlowById(userID, flowID uuid.UUID) error {
	// Service : Hard Delete Flow By Id
	err := s.flowRepo.HardDeleteFlowById(flowID, userID)
	if err != nil {
		return err
	}

	// Service : Hard Delete Flow Relation By Flow Id
	err = s.flowRelationRepo.HardDeleteFlowRelationByFlowId(flowID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}

func (s *flowService) CreateFlow(flow *models.Flow, userID uuid.UUID) (*models.Flow, error) {
	var totalAmmount int

	// Repo : Create Flow
	flow, err := s.flowRepo.CreateFlow(flow, userID)
	if err != nil {
		return nil, err
	}

	// Repo : Create Flow Relation
	var savedRelations []models.FlowRelation
	for _, dt := range flow.FlowRelations {
		var flowRelation models.FlowRelation
		flowRelation.FlowId = flow.ID
		flowRelation.PocketId = dt.PocketId
		flowRelation.Ammount = dt.Ammount

		resFlowRel, err := s.flowRelationRepo.CreateFlowRelation(&flowRelation, userID)
		if err != nil {
			return nil, err
		}

		savedRelations = append(savedRelations, *resFlowRel)
		totalAmmount += dt.Ammount
	}

	flow.FlowRelations = savedRelations

	// Repo : Create History
	historyContext := fmt.Sprintf("%s (%s) with ammount %d", flow.FlowName, flow.FlowCategory, totalAmmount)
	history := &models.History{
		HistoryType:    "Create Flow",
		HistoryContext: historyContext,
	}
	err = s.historyRepo.CreateHistory(history, userID)
	if err != nil {
		return nil, err
	}

	return flow, err
}
