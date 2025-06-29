package history

import (
	"moneh/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// History Interface
type HistoryRepository interface {
	FindMyHistory(userID uuid.UUID) ([]models.History, error)
	HardDeleteHistoryByID(ID, createdBy uuid.UUID) error
}

// History Struct
type historyRepository struct {
	db *gorm.DB
}

// History Constructor
func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}

func (r *historyRepository) FindMyHistory(userID uuid.UUID) ([]models.History, error) {
	// Model
	var histories []models.History

	// Query
	if err := r.db.Where("created_by", userID).Find(&histories).Error; err != nil {
		return nil, err
	}

	return histories, nil
}

func (r *historyRepository) HardDeleteHistoryByID(ID, createdBy uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", createdBy).Delete(&models.History{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
