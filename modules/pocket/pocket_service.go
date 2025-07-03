package pocket

import (
	"errors"
	"fmt"
	"moneh/models"
	"moneh/modules/history"
	"moneh/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Pocket Interface
type PocketService interface {
	GetAllPocket(pagination utils.Pagination, userID uuid.UUID) ([]models.Pocket, int, error)
	CreatePocket(pocket *models.Pocket, userID uuid.UUID) (*models.Pocket, error)
}

// Pocket Struct
type pocketService struct {
	pocketRepo  PocketRepository
	historyRepo history.HistoryRepository
}

// Pocket Constructor
func NewPocketService(pocketRepo PocketRepository, historyRepo history.HistoryRepository) PocketService {
	return &pocketService{
		pocketRepo:  pocketRepo,
		historyRepo: historyRepo,
	}
}

func (s *pocketService) GetAllPocket(pagination utils.Pagination, userID uuid.UUID) ([]models.Pocket, int, error) {
	return s.pocketRepo.FindAllPocket(pagination, userID)
}

func (s *pocketService) CreatePocket(pocket *models.Pocket, userID uuid.UUID) (*models.Pocket, error) {
	// Repo : Check Pokcet By Name
	isFound, err := s.pocketRepo.CheckPocketByName(pocket.PocketName, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	if isFound {
		return nil, errors.New("pocket with the same name already exists")
	}

	// Repo : Create Pocket
	pocket, err = s.pocketRepo.CreatePocket(pocket, userID)
	if err != nil {
		return nil, err
	}

	// Repo : Create History
	historyContext := fmt.Sprintf("%s (%s) with limit %d", pocket.PocketName, pocket.PocketType, pocket.PocketAmmount)
	history := &models.History{
		HistoryType:    "Create Pocket",
		HistoryContext: historyContext,
	}
	err = s.historyRepo.CreateHistory(history, userID)
	if err != nil {
		return nil, err
	}

	return pocket, err
}
