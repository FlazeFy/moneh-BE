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
	DeleteHistoryById(id uuid.UUID) error
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

func (s *historyService) DeleteHistoryById(id uuid.UUID) error {
	// Repo : Get My History
	affected_row, err := s.historyRepo.DeleteById(id)
	if err != nil {
		return err
	}
	if affected_row == 0 {
		return errors.New("history not found")
	}

	return nil
}
