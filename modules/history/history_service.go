package history

import (
	"errors"
	"moneh/models"
	"moneh/utils"

	"github.com/google/uuid"
)

type HistoryService interface {
	GetAllHistory(pagination utils.Pagination) ([]models.AllHistory, int64, error)
	GetMyHistory(pagination utils.Pagination, id uuid.UUID, typeUser string) ([]models.History, int64, error)
}

type historyService struct {
	historyRepo HistoryRepository
}

func NewHistoryService(historyRepo HistoryRepository) HistoryService {
	return &historyService{
		historyRepo: historyRepo,
	}
}

func (s *historyService) GetAllHistory(pagination utils.Pagination) ([]models.AllHistory, int64, error) {
	// Repo : Get All History
	history, total, err := s.historyRepo.FindAll(pagination)
	if err != nil {
		return nil, 0, err
	}
	if history == nil {
		return nil, 0, errors.New("history not found")
	}

	return history, total, nil
}

func (s *historyService) GetMyHistory(pagination utils.Pagination, id uuid.UUID, typeUser string) ([]models.History, int64, error) {
	// Repo : Get My History
	history, total, err := s.historyRepo.FindMy(pagination, id, typeUser)
	if err != nil {
		return nil, 0, err
	}
	if history == nil {
		return nil, 0, errors.New("history not found")
	}

	return history, total, nil
}
