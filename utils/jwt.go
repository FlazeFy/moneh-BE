package utils

import (
	"moneh/config"
	"moneh/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateToken(userId uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId.String(),
		"exp":     time.Now().Add(config.GetJWTExpirationDuration()).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.GetJWTSecret())
}

func HashPassword(user models.User, password string) (*models.User, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPass)

	return &user, nil
}

func CheckPassword(account models.Account, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(account.GetPassword()), []byte(password))
}
