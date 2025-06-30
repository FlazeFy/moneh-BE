package seeders

import (
	"log"
	"moneh/config"
	"moneh/factories"
	"moneh/modules/dictionary"
)

func SeedDictionaries(repo dictionary.DictionaryRepository) {
	// Empty Table
	repo.DeleteAll()

	var seedData = []struct {
		DictionaryType  string
		DictionaryNames []string
	}{
		{"currency", config.Currencies},
		{"flow_type", config.FlowTypes},
		{"flow_category", config.FlowCategories},
		{"pocket_type", config.PocketTypes},
	}

	// Fill Table
	var success = 0
	for _, dt := range seedData {
		for _, dictionaryName := range dt.DictionaryNames {
			dct := factories.DictionaryFactory(dictionaryName, dt.DictionaryType)
			err := repo.CreateDictionary(&dct)
			if err != nil {
				log.Printf("failed to seed dictionary %s/%s: %v\n", dt.DictionaryType, dictionaryName, err)
			}
			success++
		}
	}
	log.Printf("Seeder : Success to seed %d Dictionary", success)
}
