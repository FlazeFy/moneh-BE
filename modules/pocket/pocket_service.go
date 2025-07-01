package pocket

import (
	"moneh/models"
	"moneh/utils"

	"github.com/google/uuid"
)

// Pocket Interface
type PocketService interface {
	GetAllPocket(pagination utils.Pagination, userID uuid.UUID) ([]models.Pocket, int, error)
}

// Pocket Struct
type pocketService struct {
	pocketRepo PocketRepository
}

// Pocket Constructor
func NewPocketService(pocketRepo PocketRepository) PocketService {
	return &pocketService{
		pocketRepo: pocketRepo,
	}
}

func (r *pocketService) GetAllPocket(pagination utils.Pagination, userID uuid.UUID) ([]models.Pocket, int, error) {
	return r.pocketRepo.FindAllPocket(pagination, userID)
}
