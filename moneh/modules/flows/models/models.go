package models

type (
	GetFlow struct {
		Id            string `json:"id"`
		FlowsType     string `json:"flows_type"`
		FlowsCategory string `json:"flows_category"`
		FlowsName     string `json:"flows_name"`
		FlowsDesc     string `json:"flows_desc"`
		FlowsAmmount  int    `json:"flows_ammount"`
		FlowsTag      string `json:"flows_tag"`
		IsShared      string `json:"is_shared"`
	}
	GetSummaryByType struct {
		Average      int `json:"average"`
		TotalItem    int `json:"total_item"`
		TotalAmmount int `json:"total_ammount"`
	}
)
