package seeders

import (
	"log"
	"moneh/factories/dummies"
	"moneh/modules/systems/models"
	"moneh/modules/systems/repositories"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"

	"github.com/bxcodec/faker/v3"
)

func SeedDictionaries(total int, showRes bool) {
	var obj models.PostDictionaryByType
	idx := 0
	var logs string

	for idx < total {
		// Data
		obj.DctType = generator.GetSlug(dummies.DummyDctType())
		obj.DctName = faker.Word()

		result, err := repositories.PostDictionary(obj)
		if err != nil {
			log.Println(err.Error())
		}

		if showRes {
			log.Println(result.Data)
			if strData, ok := result.Data.(string); ok {
				logs += strData + "\n"
			}
		}
		idx++
	}

	if showRes {
		response.ResponsePrinter("txt", "seeder_dictionaries", logs)
	}
}
