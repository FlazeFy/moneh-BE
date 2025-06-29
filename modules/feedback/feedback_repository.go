package feedback

import (
	"moneh/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Feedback Interface
type FeedbackRepository interface {
	CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error
	FindAllFeedback() ([]models.Feedback, error)

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

func (r *feedbackRepository) FindAllFeedback() ([]models.Feedback, error) {
	// Model
	var feedbacks []models.Feedback

	// Query
	if err := r.db.Preload("User").Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	if len(feedbacks) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return feedbacks, nil
}

func (r *feedbackRepository) CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error {
	// Default
	feedback.ID = uuid.New()
	feedback.CreatedAt = time.Now()
	feedback.CreatedBy = userID

	// Query
	return r.db.Create(feedback).Error
}

// For Seeder
func (r *feedbackRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Feedback{}).Error
}
