package models

type (
	GetMyProfile struct {
		Username  string `json:"username"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		ImageUrl  string `json:"image_url"`

		// Props
		AcceptedAt string `json:"accepted_at"`
	}
)
