package factories

import (
	"encoding/json"
	"moneh/config"
	"moneh/models"
	"moneh/utils"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

func FlowFactory() models.Flow {
	flowDesc := gofakeit.LoremIpsumSentence(gofakeit.Number(2, 5))

	var flowTag []byte
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
		flowTag = jsonData
	} else {
		flowTag = nil
	}

	return models.Flow{
		FlowType:      gofakeit.RandomString(config.FlowTypes),
		FlowCategory:  gofakeit.RandomString(config.FlowCategories),
		FlowName:      gofakeit.LoremIpsumSentence(gofakeit.Number(2, 5)),
		FlowDesc:      &flowDesc,
		FlowTag:       flowTag,
		IsSplitBill:   gofakeit.Bool(),
		IsMultiPocket: gofakeit.Bool(),
	}
}
