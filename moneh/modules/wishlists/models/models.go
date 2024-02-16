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
		IsAchieved       int    `json:"is_achieved"`
	}
	GetSummary struct {
		Average           int    `json:"average"`
		Achieved          int    `json:"achieved"`
		TotalItem         int    `json:"total_item"`
		TotalAmmount      int    `json:"total_ammount"`
		MostExpensive     int    `json:"most_expensive"`
		Cheapest          int    `json:"cheapest"`
		MostType          string `json:"most_type"`
		MostExpensiveName string `json:"most_expensive_name"`
		CheapestName      string `json:"cheapest_name"`
	}
	PostWishlist struct {
		WishlistName     string `json:"wishlists_name"`
		WishlistDesc     string `json:"wishlists_desc"`
		WishlistImgUrl   string `json:"wishlists_img_url"`
		WishlistType     string `json:"wishlists_type"`
		WishlistPriority string `json:"wishlists_priority"`
		WishlistPrice    int    `json:"wishlists_price"`
		IsAchieved       int    `json:"is_achieved"`

		// Properties
		CreatedAt string `json:"created_at"`
	}
)
