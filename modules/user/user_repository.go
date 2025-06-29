package user

import (
	"errors"
	"fmt"
	"moneh/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsernameOrEmail(username, email string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindById(id, role string) (*models.MyProfile, error)
	CreateUser(user *models.User) (*models.User, error)

	// For Seeder
	DeleteAll() error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsernameOrEmail(username, email string) (*models.User, error) {
	// Models
	var user models.User

	// Query
	err := r.db.Where("username = ? OR email = ?", username, email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	// Models
	var user models.User

	// Query
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *userRepository) FindById(id, role string) (*models.MyProfile, error) {
	// Models
	var user models.MyProfile
	var tableName = fmt.Sprintf("%ss", role)
	if role == "guest" {
		tableName = "users"
	}

	// Query
	err := r.db.Table(tableName).
		Select("username, email, telegram_is_valid, telegram_user_id, created_at").
		Where("id = ?", id).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.TelegramIsValid = false

	// Query
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// For Seeder
func (r *userRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.User{}).Error
}
