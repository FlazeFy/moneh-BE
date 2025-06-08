package auth

import (
	"errors"
	"moneh/models"
	"moneh/utils"
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(user *models.User) (string, error)
	Login(email, password string) (string, string, error)
}

type authService struct {
	userRepo UserRepository
}

func NewAuthService(userRepo UserRepository) AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Register(user *models.User) (string, error) {
	// Check duplicate
	existing, err := s.userRepo.FindByUsernameOrEmail(user.Username, user.Email)
	if err != nil {
		return "", err
	}
	if existing != nil {
		return "", errors.New("username or email has already been used")
	}

	// Utils : Hash Password
	if err := utils.HashPassword(user, user.Password); err != nil {
		return "", err
	}

	// Mapping
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.TelegramIsValid = false

	// Repo : Create Register
	if err := s.userRepo.Create(user); err != nil {
		return "", err
	}

	// Utils : Generate Token
	token, err := utils.GenerateToken(user.ID, "user")
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) Login(email, password string) (string, string, error) {
	// Model
	var account models.Account
	var role string

	// Repo : Check Admin By Email
	admin, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", "", err
	}
	if admin != nil {
		account = admin
		role = "admin"
	}

	// Repo : Check User (Guest) By Email
	if account == nil {
		user, err := s.userRepo.FindByEmail(email)
		if err != nil {
			return "", "", err
		}
		if user != nil {
			account = user
			role = "guest"
		}
	}

	if account == nil {
		return "", "", errors.New("account not found")
	}

	// Utils : Compare Password
	if err := utils.CheckPassword(account, password); err != nil {
		return "", "", errors.New("invalid password")
	}

	// Utils : Generate Token
	token, err := utils.GenerateToken(account.GetID(), role)
	if err != nil {
		return "", "", err
	}

	return token, role, nil
}
