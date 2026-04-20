package history

import (
	"moneh/models"
	"moneh/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// History Interface
type HistoryRepository interface {
	FindMyHistory(userID uuid.UUID, pagination utils.Pagination) ([]models.AllHistory, int, error)
	HardDeleteHistoryByID(ID, createdBy uuid.UUID) error

	// For Seeder
	CreateHistory(history *models.History, userID uuid.UUID) error
	DeleteAll() error
}

// History Struct
type historyRepository struct {
	db *gorm.DB
}

// History Constructor
func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}

func (r *historyRepository) FindMyHistory(userID uuid.UUID, pagination utils.Pagination) ([]models.AllHistory, int, error) {
	// Model
	var total int
	var histories []models.AllHistory

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit

	// Query
	err := r.db.Table("histories").
		Select("id", "history_type", "history_context", "created_at").
		Order("created_at ASC").
		Where("created_by", userID).
		Limit(pagination.Limit).
		Offset(offset).
		Find(&histories).Error

	if err != nil {
		return nil, 0, err
	}

	total = len(histories)
	return histories, total, nil
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

// For Seeder
func (r *historyRepository) CreateHistory(history *models.History, userID uuid.UUID) error {
	// Default
	history.ID = uuid.New()
	history.CreatedAt = time.Now()
	history.CreatedBy = userID

	// Query
	return r.db.Create(history).Error
}

func (r *historyRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.History{}).Error
}
