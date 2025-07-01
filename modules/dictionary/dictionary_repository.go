package dictionary

import (
	"moneh/models"
	"moneh/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Dictionary Interface
type DictionaryRepository interface {
	CreateDictionary(dictionary *models.Dictionary) error
	FindAllDictionary(pagination utils.Pagination) ([]models.Dictionary, int, error)

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

func (r *dictionaryRepository) FindAllDictionary(pagination utils.Pagination) ([]models.Dictionary, int, error) {
	// Model
	var total int
	var dictionaries []models.Dictionary

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit

	// Query
	err := r.db.Order("dictionary_type ASC").
		Order("dictionary_name ASC").
		Limit(pagination.Limit).
		Offset(offset).
		Find(&dictionaries).Error

	if err != nil {
		return nil, 0, err
	}

	total = len(dictionaries)
	return dictionaries, total, nil
}

// For Seeder
func (r *dictionaryRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Dictionary{}).Error
}
