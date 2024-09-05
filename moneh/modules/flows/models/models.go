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
		IsShared      int    `json:"is_shared"`
		CreatedAt     string `json:"created_at"`
		UpdatedAt     string `json:"updated_at"`
	}
	GetFlowExport struct {
		FlowsType     string `json:"flows_type"`
		FlowsCategory string `json:"flows_category"`
		FlowsName     string `json:"flows_name"`
		FlowsDesc     string `json:"flows_desc"`
		FlowsAmmount  int    `json:"flows_ammount"`
		CreatedAt     string `json:"created_at"`
	}
	GetSummaryByType struct {
		Average      int `json:"average"`
		TotalItem    int `json:"total_item"`
		TotalAmmount int `json:"total_ammount"`
	}
	GetTotalItemAmmountPerDateByType struct {
		TotalItem    int `json:"total_item"`
		TotalAmmount int `json:"total_ammount"`

		// Properties
		Context string `json:"context"`
	}
	GetMonthlyFlow struct {
		Title   string `json:"title"`
		Context string `json:"context"`
	}
	GetMonthlyFlowTotal struct {
		Title   string `json:"title"`
		Type    string `json:"type"`
		Context string `json:"context"`
	}
)
