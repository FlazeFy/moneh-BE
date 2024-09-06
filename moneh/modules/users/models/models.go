package models

type (
	GetMyProfile struct {
		ID              string `json:"id"`
		Username        string `json:"username"`
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		Email           string `json:"email"`
		ImageUrl        string `json:"image_url"`
		TelegramUserId  string `json:"telegram_user_id"`
		TelegramIsValid int    `json:"telegram_is_valid"`

		// Props
		AcceptedAt string `json:"accepted_at"`
	}
)
