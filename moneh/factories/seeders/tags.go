package seeders

import (
	"fmt"
	"moneh/modules/systems/models"
	"moneh/modules/systems/repositories"
	"moneh/packages/helpers/generator"

	"github.com/bxcodec/faker/v3"
)

func SeedTags(total int, showRes bool) {
	var obj models.PostTag
	idx := 0

	for idx < total {
		// Data
		name := faker.Word()
		obj.TagSlug = generator.GetSlug(name)
		obj.TagName = name

		result, err := repositories.PostTag(obj)
		if err != nil {
			fmt.Println(err.Error())
		}

		if showRes {
			fmt.Println(result.Data)
		}
		idx++
	}
}
