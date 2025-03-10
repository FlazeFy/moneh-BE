package seeders

import (
	"fmt"
	"math/rand"
	"moneh/factories/dummies"
	"moneh/modules/wishlists/models"
	"moneh/modules/wishlists/repositories"
	"moneh/packages/helpers/generator"
	"moneh/packages/helpers/response"
	"time"

	"github.com/bxcodec/faker/v3"
)

func SeedWishlists(total int, showRes bool) {
	rand.Seed(time.Now().UnixNano())
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjY3NzQ1NDc5MTIsImxldmVsIjoiYXBwbGljYXRpb24iLCJ1c2VybmFtZSI6InRlc3RlcnVzZXIifQ.pJv9kjPbUp78-0McyOCJyB9raL0V2nR-jjYVnKlT_7s"

	var obj models.PostWishlist
	idx := 0
	var logs string

	for idx < total {
		// Data
		obj.WishlistName = faker.Word()
		obj.WishlistDesc = faker.Paragraph()
		obj.WishlistImgUrl = faker.URL()
		obj.WishlistType = dummies.DummyWishlistType()
		obj.WishlistPrice = generator.GeneratePrice(10000000, 1000)
		obj.WishlistPriority = dummies.DummyPriority()
		obj.IsAchieved = int(rand.Float64())

		result, err := repositories.PostWishlist(obj, token)
		if err != nil {
			fmt.Println(err.Error())
		}

		if showRes {
			fmt.Println(result.Data)
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
