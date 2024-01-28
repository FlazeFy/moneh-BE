package seeders

import (
	"fmt"
	"math/rand"
	"moneh/factories/dummies"
	"moneh/modules/wishlists/models"
	"moneh/modules/wishlists/repositories"
	"time"

	"github.com/bxcodec/faker/v3"
)

func SeedWishlists(total int, showRes bool) {
	rand.Seed(time.Now().UnixNano())

	var obj models.PostWishlist
	idx := 0

	for idx < total {
		// Data
		obj.WishlistName = faker.Word()
		obj.WishlistDesc = faker.Paragraph()
		obj.WishlistImgUrl = faker.URL()
		obj.WishlistType = dummies.DummyWishlistType()
		obj.WishlistPriority = dummies.DummyPriority()
		obj.IsAchieved = int(rand.Float64())

		result, err := repositories.PostWishlist(obj)
		if err != nil {
			fmt.Println(err.Error())
		}

		if showRes {
			fmt.Println(result.Data)
		}
		idx++
	}
}
