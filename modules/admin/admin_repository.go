package admin

import (
	"errors"
	"moneh/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByEmail(email string) (*models.Admin, error)
	FindAllAdminContact() ([]models.UserContact, error)

	// For Seeder
	CreateAdmin(admin *models.Admin) error
	DeleteAll() error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) FindByEmail(email string) (*models.Admin, error) {
	// Models
	var admin models.Admin

	// Query
	err := r.db.Where("email = ?", email).First(&admin).Error
	if err != nil {
		return nil, err
	}

	return &admin, nil
}

// For Task Scheduler
func (r *adminRepository) FindAllAdminContact() ([]models.UserContact, error) {
	// Model
	var contact []models.UserContact

	// Query
	result := r.db.Table("admins").
		Select("username, email, telegram_user_id, telegram_is_valid").
		Find(&contact)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || len(contact) == 0 {
		return nil, errors.New("admin contact not found")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return contact, nil
}

// For Seeder
func (r *adminRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Admin{}).Error
}
func (r *adminRepository) CreateAdmin(admin *models.Admin) error {
	admin.ID = uuid.New()
	admin.CreatedAt = time.Now()
	admin.TelegramIsValid = false

	// Query
	return r.db.Create(admin).Error
}
