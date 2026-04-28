package feedback

import (
	"moneh/models"
	"moneh/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Feedback Interface
type FeedbackRepository interface {
	CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error
	FindAllFeedback(pagination utils.Pagination) ([]models.AllFeedbackData, int, error)
	HardDeleteFeedbackByID(ID uuid.UUID) error

	// For Seeder
	DeleteAll() error
}

// Feedback Struct
type feedbackRepository struct {
	db *gorm.DB
}

// Feedback Constructor
func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{db: db}
}

func (r *feedbackRepository) FindAllFeedback(pagination utils.Pagination) ([]models.AllFeedbackData, int, error) {
	// Model
	var total int
	var feedbacks []models.AllFeedbackData

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit

	// Query
	err := r.db.Table("feedbacks").
		Select(`feedbacks.id, feedbacks.feedback_body, feedbacks.feedback_rate, feedbacks.created_at, users.username, users.email`).
		Joins("LEFT JOIN users ON users.id = feedbacks.created_by").
		Order("feedbacks.created_at ASC").
		Limit(pagination.Limit).
		Offset(offset).
		Scan(&feedbacks).Error

	if err != nil {
		return nil, 0, err
	}

	total = len(feedbacks)
	return feedbacks, total, nil
}

func (r *feedbackRepository) CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error {
	// Default
	feedback.ID = uuid.New()
	feedback.CreatedAt = time.Now()
	feedback.CreatedBy = userID

	// Query
	return r.db.Create(feedback).Error
}

func (r *feedbackRepository) HardDeleteFeedbackByID(ID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Delete(&models.Feedback{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// For Seeder
func (r *feedbackRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Feedback{}).Error
}
