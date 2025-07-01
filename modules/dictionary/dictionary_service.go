package dictionary

import (
	"moneh/models"
	"moneh/utils"
)

// Dictionary Interface
type DictionaryService interface {
	GetAllDictionary(pagination utils.Pagination) ([]models.Dictionary, int, error)
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
