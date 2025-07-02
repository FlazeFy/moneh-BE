package dictionary

import (
	"errors"
	"moneh/models"
	"moneh/utils"

	"gorm.io/gorm"
)

// Dictionary Interface
type DictionaryService interface {
	GetAllDictionary(pagination utils.Pagination) ([]models.Dictionary, int, error)
	CreateDictionary(dictionary *models.Dictionary) error
}

// Dictionary Struct
type dictionaryService struct {
	dictionaryRepo DictionaryRepository
}

// Dictionary Constructor
func NewDictionaryService(dictionaryRepo DictionaryRepository) DictionaryService {
	return &dictionaryService{
		dictionaryRepo: dictionaryRepo,
	}
}

func (r *dictionaryService) GetAllDictionary(pagination utils.Pagination) ([]models.Dictionary, int, error) {
	return r.dictionaryRepo.FindAllDictionary(pagination)
}

func (r *dictionaryService) CreateDictionary(dictionary *models.Dictionary) error {
	// Repo : Find Dictionary By Type
	_, err := r.dictionaryRepo.FindOneDictionaryByName(dictionary.DictionaryName)
	if err == nil {
		return errors.New("dictionary already exist")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return r.dictionaryRepo.CreateDictionary(dictionary)
}
