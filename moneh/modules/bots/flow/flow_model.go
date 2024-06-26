package flow

type (
	GetAllFlowModel struct {
		FlowsType     string `json:"flows_type"`
		FlowsCategory string `json:"flows_category"`
		FlowsName     string `json:"flows_name"`
		FlowsAmmount  int    `json:"flows_ammount"`
		CreatedAt     string `json:"created_at"`
	}
	GetFlowDaily struct {
		Context       string `json:"context"`
		TotalSpending int    `json:"total_spending"`
		TotalIncome   int    `json:"total_income"`
	}
)
