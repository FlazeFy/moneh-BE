package pocket

import (
	"moneh/models"
	"moneh/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Pocket Interface
type PocketRepository interface {
	CreatePocket(pocket *models.Pocket, userID uuid.UUID) (*models.Pocket, error)
	FindAllPocket(pagination utils.Pagination, userID uuid.UUID) ([]models.Pocket, int, error)
	CheckPocketByName(pocketName string, userID uuid.UUID) (bool, error)

	// For Seeder
	DeleteAll() error
	FindOneRandomByUserID(userID uuid.UUID) (*models.Pocket, error)
}

// Pocket Struct
type pocketRepository struct {
	db *gorm.DB
}

// Pocket Constructor
func NewPocketRepository(db *gorm.DB) PocketRepository {
	return &pocketRepository{db: db}
}

func (r *pocketRepository) CreatePocket(pocket *models.Pocket, userID uuid.UUID) (*models.Pocket, error) {
	// Default
	pocket.ID = uuid.New()
	pocket.CreatedAt = time.Now()
	pocket.CreatedBy = userID
	pocket.UpdatedAt = nil

	// Query
	if err := r.db.Create(pocket).Error; err != nil {
		return nil, err
	}

	return pocket, nil
}

func (r *pocketRepository) FindAllPocket(pagination utils.Pagination, userID uuid.UUID) ([]models.Pocket, int, error) {
	// Model
	var total int
	var pockets []models.Pocket

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit

	// Query
	err := r.db.Where("created_by", userID).
		Order("pocket_ammount ASC").
		Order("pocket_name ASC").
		Limit(pagination.Limit).
		Offset(offset).
		Find(&pockets).Error

	if err != nil {
		return nil, 0, err
	}

	total = len(pockets)
	return pockets, total, nil
}

func (r *pocketRepository) CheckPocketByName(pocketName string, userID uuid.UUID) (bool, error) {
	// Model
	var pocket models.Pocket

	// Query
	result := r.db.Unscoped().Where("LOWER(pocket_name) = LOWER(?) AND created_by = ?", pocketName, userID).First(&pocket)

	// Response
	if result.Error != nil {
		return false, result.Error
	}
	if pocket.ID == uuid.Nil {
		return false, nil
	}

	return true, nil
}

// For Seeder
func (r *pocketRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Pocket{}).Error
}

func (r *pocketRepository) FindOneRandomByUserID(userID uuid.UUID) (*models.Pocket, error) {
	var pocket models.Pocket

	err := r.db.Where("created_by", userID).
		Where("pocket_limit < pocket_ammount").
		Order("RANDOM()").
		Limit(1).
		First(&pocket).Error

	if err != nil {
		return nil, err
	}

	return &pocket, nil
}
