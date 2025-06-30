package factories

import (
	"encoding/json"
	"moneh/config"
	"moneh/models"
	"moneh/utils"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

func PocketFactory() models.Pocket {
	pocketDesc := gofakeit.LoremIpsumSentence(gofakeit.Number(2, 5))

	var pocketTag []byte
	if gofakeit.Bool() {
		numTags := gofakeit.Number(1, 5)
		var tags []models.FlowPocketTags

		for i := 0; i < numTags; i++ {
			tagName := utils.Capitalize(gofakeit.VerbAction())
			tagSlug := strings.ToLower(strings.ReplaceAll(tagName, " ", "_"))

			tags = append(tags, models.FlowPocketTags{
				TagName: tagName,
				TagSlug: tagSlug,
			})
		}

		jsonData, _ := json.Marshal(tags)
		pocketTag = jsonData
	} else {
		pocketTag = nil
	}

	var pocketLimit *int
	if gofakeit.Bool() {
		limit := gofakeit.Number(1, 25) * 100000
		pocketLimit = &limit
	} else {
		pocketLimit = nil
	}

	return models.Pocket{
		PocketType:    gofakeit.RandomString(config.PocketTypes),
		PocketName:    gofakeit.LoremIpsumSentence(gofakeit.Number(2, 3)),
		PocketDesc:    &pocketDesc,
		PocketAmmount: gofakeit.Number(1, 25) * 1000000,
		PocketLimit:   pocketLimit,
		PocketTags:    pocketTag,
	}
}
