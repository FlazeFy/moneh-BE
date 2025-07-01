package user

import (
	"moneh/models"

	"github.com/google/uuid"
)

// User Interface
type UserService interface {
	GetMyProfile(userID uuid.UUID) (*models.MyProfile, error)
}

// User Struct
type userService struct {
	userRepo UserRepository
}

// User Constructor
func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (r *userService) GetMyProfile(userID uuid.UUID) (*models.MyProfile, error) {
	// Repo : Find User By User Id
	return r.userRepo.FindById(userID)
}
