package models

type (
	GetWishlistHeaders struct {
		Id               string `json:"id"`
		WishlistName     string `json:"wishlists_name"`
		WishlistDesc     string `json:"wishlists_desc"`
		WishlistImgUrl   string `json:"wishlists_img_url"`
		WishlistType     string `json:"wishlists_type"`
		WishlistPriority string `json:"wishlists_priority"`
	}
)
