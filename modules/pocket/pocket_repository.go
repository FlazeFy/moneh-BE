package pocket

import (
	"moneh/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Pocket Interface
type PocketRepository interface {
	CreatePocket(pocket *models.Pocket, userID uuid.UUID) error

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

func (r *pocketRepository) CreatePocket(pocket *models.Pocket, userID uuid.UUID) error {
	// Default
	pocket.ID = uuid.New()
	pocket.CreatedAt = time.Now()
	pocket.CreatedBy = userID
	pocket.UpdatedAt = nil

	// Query
	return r.db.Create(pocket).Error
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
