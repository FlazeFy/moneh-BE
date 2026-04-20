package history

import (
	"moneh/models"
	"moneh/utils"

	"github.com/google/uuid"
)

// History Interface
type HistoryService interface {
	GetMyHistory(userID uuid.UUID, pagination utils.Pagination) ([]models.AllHistory, int, error)
	HardDeleteHistoryByID(ID, createdBy uuid.UUID) error
}

// History Struct
type historyService struct {
	historyRepo HistoryRepository
}

// History Constructor
func NewHistoryService(historyRepo HistoryRepository) HistoryService {
	return &historyService{
		historyRepo: historyRepo,
	}
}

func (r *historyService) GetMyHistory(userID uuid.UUID, pagination utils.Pagination) ([]models.AllHistory, int, error) {
	return r.historyRepo.FindMyHistory(userID, pagination)
}

func (r *historyService) HardDeleteHistoryByID(ID, createdBy uuid.UUID) error {
	return r.historyRepo.HardDeleteHistoryByID(ID, createdBy)
}
