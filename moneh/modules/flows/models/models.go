package models

type (
	GetAnimalHeaders struct {
		Id            string `json:"id"`
		FlowsType     string `json:"flows_type"`
		FlowsCategory string `json:"flows_category"`
		FlowsName     string `json:"flows_name"`
		FlowsDesc     string `json:"flows_desc"`
		FlowsAmmount  string `json:"flows_ammount"`
		FlowsTag      string `json:"flows_tag"`
		IsShared      string `json:"is_shared"`
	}
)
