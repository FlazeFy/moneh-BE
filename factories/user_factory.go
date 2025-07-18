package factories

import (
	"moneh/config"
	"moneh/models"

	"github.com/brianvoe/gofakeit/v6"
	"golang.org/x/crypto/bcrypt"
)

func UserFactory(username, email, telegramUserId, password *string, isValid bool) models.User {
	var finalUsername string
	if username != nil && *username != "" {
		finalUsername = *username
	} else {
		finalUsername = gofakeit.Username()
	}

	var finalEmail string
	if email != nil && *email != "" {
		finalEmail = *email
	} else {
		finalEmail = gofakeit.Email()
	}

	var pwd string
	if password != nil && *password != "" {
		pwd = *password
	} else {
		pwd = "nopass123"
	}
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	var finalTelegramUserId *string
	if telegramUserId != nil && *telegramUserId != "" {
		finalTelegramUserId = telegramUserId
	} else {
		finalTelegramUserId = nil
	}

	return models.User{
		Username:        finalUsername,
		Password:        string(hashedPass),
		TelegramUserId:  finalTelegramUserId,
		TelegramIsValid: isValid,
		Email:           finalEmail,
		Currency:        gofakeit.RandomString(config.Currencies),
	}
}
