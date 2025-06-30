package dictionary

import (
	"moneh/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Dictionary Interface
type DictionaryRepository interface {
	CreateDictionary(dictionary *models.Dictionary) error

	// For Seeder
	DeleteAll() error
}

// Dictionary Struct
type dictionaryRepository struct {
	db *gorm.DB
}

// Dictionary Constructor
func NewDictionaryRepository(db *gorm.DB) DictionaryRepository {
	return &dictionaryRepository{db: db}
}

func (r *dictionaryRepository) CreateDictionary(dictionary *models.Dictionary) error {
	// Default
	dictionary.ID = uuid.New()
	dictionary.CreatedAt = time.Now()

	// Query
	return r.db.Create(dictionary).Error
}

// For Seeder
func (r *dictionaryRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Dictionary{}).Error
}
