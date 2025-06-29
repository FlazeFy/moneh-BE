package admin

import (
	"moneh/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AdminRepository interface {
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
