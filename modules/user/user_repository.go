package user

import (
	"moneh/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsernameOrEmail(username, email string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindById(ID uuid.UUID) (*models.MyProfile, error)
	CreateUser(user *models.User) (*models.User, error)

	// For Seeder
	DeleteAll() error
	FindOneRandom() (*models.User, error)
	FindOneHasFlowAndPocketRandom() (*models.User, error)
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
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	// Models
	var user models.User

	// Query
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (r *userRepository) FindById(ID uuid.UUID) (*models.MyProfile, error) {
	// Models
	var user models.MyProfile

	// Query
	err := r.db.Table("users").
		Select("username, email, telegram_is_valid, telegram_user_id, currency, created_at").
		Where("id = ?", ID).
		First(&user).Error

	if err != nil {
		return nil, err
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

func (r *userRepository) FindOneRandom() (*models.User, error) {
	var user models.User

	err := r.db.Order("RANDOM()").Limit(1).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindOneHasFlowAndPocketRandom() (*models.User, error) {
	var users models.User

	err := r.db.Select("users.id, users.username, users.password, users.email, users.telegram_user_id, users.telegram_is_valid, users.created_at").
		Joins("JOIN flows f ON f.created_by = users.id").
		Joins("JOIN pockets p ON p.created_by = users.id").
		Group("users.id").
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return &users, nil
}
