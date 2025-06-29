package factories

import (
	"moneh/config"
	"moneh/models"

	"github.com/brianvoe/gofakeit/v6"
	"golang.org/x/crypto/bcrypt"
)

func UserFactory() models.User {
	password := "nopass123"
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return models.User{
		Username:        gofakeit.Username(),
		Password:        string(hashedPass),
		TelegramUserId:  nil,
		TelegramIsValid: false,
		Email:           gofakeit.Email(),
		Currency:        gofakeit.RandomString(config.Currencies),
	}
}
