package models

type (
	GetFeedbacks struct {
		FdbRate string `json:"feedbacks_rate"`
		FdbDesc string `json:"feedbacks_desc"`

		// Props
		CreatedAt string `json:"created_at"`
	}
)
