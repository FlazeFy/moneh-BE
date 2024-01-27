package models

type (
	GetDictionaryByType struct {
		ID      int    `json:"id"`
		DctName string `json:"dictionaries_name"`
	}
	PostDictionaryByType struct {
		DctType string `json:"dictionaries_type"`
		DctName string `json:"dictionaries_name"`
	}
)
