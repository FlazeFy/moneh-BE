package models

type (
	GetFeedbacks struct {
		FdbRate int    `json:"feedbacks_rate"`
		FdbDesc string `json:"feedbacks_desc"`

		// Props
		CreatedAt string `json:"created_at"`
	}
	PostFeedback struct {
		FdbRate int    `json:"feedbacks_rate"`
		FdbDesc string `json:"feedbacks_desc"`
	}
)
