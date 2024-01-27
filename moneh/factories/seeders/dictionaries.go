package seeders

import (
	"fmt"
	"moneh/factories/dummies"
	"moneh/modules/systems/models"
	"moneh/modules/systems/repositories"

	"github.com/bxcodec/faker/v3"
)

func SeedDictionaries(total int, showRes bool) {
	var obj models.PostDictionaryByType
	idx := 0

	for idx < total {
		// Data
		obj.DctType = dummies.DummyDctType()
		obj.DctName = faker.Word()

		result, err := repositories.PostDictionary(obj)
		if err != nil {
			fmt.Println(err.Error())
		}

		if showRes {
			fmt.Println(result.Data)
		}
		idx++
	}
}
