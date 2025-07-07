package errors

import (
	"moneh/models"
	"moneh/utils"
)

// Error Interface
type ErrorService interface {
	GetAllError(pagination utils.Pagination) ([]models.ErrorAudit, int64, error)
}

// Error Struct
type errorService struct {
	errorRepo ErrorRepository
}

// Error Constructor
func NewErrorService(errorRepo ErrorRepository) ErrorService {
	return &errorService{
		errorRepo: errorRepo,
	}
}

func (s *errorService) GetAllError(pagination utils.Pagination) ([]models.ErrorAudit, int64, error) {
	return s.errorRepo.FindAllError(pagination)
}
