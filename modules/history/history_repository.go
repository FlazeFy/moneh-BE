package history

import (
	"errors"
	"fmt"
	"moneh/models"
	"moneh/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type HistoryRepository interface {
	FindAll(pagination utils.Pagination) ([]models.AllHistory, int64, error)
	FindMy(pagination utils.Pagination, id uuid.UUID, typeUser string) ([]models.History, int64, error)
}

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}

func (r *historyRepository) FindAll(pagination utils.Pagination) ([]models.AllHistory, int64, error) {
	var total int64

	// Models
	var history []models.AllHistory

	// Pagination
	offset := (pagination.Page - 1) * pagination.Limit
	r.db.Model(&models.History{}).Count(&total)

	// Query
	err := r.db.Table("histories").
		Preload("User").
		Preload("Admin").
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset).
		Find(&history).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, total, nil
	}

	return history, total, nil
}

func (r *historyRepository) FindMy(pagination utils.Pagination, id uuid.UUID, typeUser string) ([]models.History, int64, error) {
	var total int64

	// Models
	var history []models.History

	// Query
	var targetCol string
	if typeUser == "admin" {
		targetCol = "admin_id"
	} else if typeUser == "user" {
		targetCol = "user_id"
	}

	// Pagination
	offset := (pagination.Page - 1) * pagination.Limit
	r.db.Model(&models.History{}).
		Where(fmt.Sprintf("%s = ?", targetCol), id).
		Where("type_user = ?", typeUser).
		Count(&total)

	err := r.db.Where(fmt.Sprintf("%s = ?", targetCol), id).
		Where("type_user = ?", typeUser).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset).
		Find(&history).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	}

	return history, total, err
}
