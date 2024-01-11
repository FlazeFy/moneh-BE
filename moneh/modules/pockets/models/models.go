package models

type (
	GetPocketHeaders struct {
		Id           string `json:"id"`
		PocketsName  string `json:"pockets_name"`
		PocketsDesc  string `json:"pockets_desc"`
		PocketsType  string `json:"pockets_type"`
		PocketsLimit int    `json:"pockets_limit"`
	}
)
