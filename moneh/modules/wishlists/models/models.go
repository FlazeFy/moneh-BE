package models

type (
	GetWishlistHeaders struct {
		Id               string `json:"id"`
		WishlistName     string `json:"wishlists_name"`
		WishlistDesc     string `json:"wishlists_desc"`
		WishlistImgUrl   string `json:"wishlists_img_url"`
		WishlistType     string `json:"wishlists_type"`
		WishlistPriority string `json:"wishlists_priority"`
		WishlistPrice    int    `json:"wishlists_price"`
	}
	GetSummary struct {
		Average       int    `json:"average"`
		Achieved      int    `json:"achieved"`
		TotalItem     int    `json:"total_item"`
		TotalAmmount  int    `json:"total_ammount"`
		MostExpensive string `json:"most_expensive"`
		Cheapest      string `json:"cheapest"`
		MostType      string `json:"most_type"`
	}
)
