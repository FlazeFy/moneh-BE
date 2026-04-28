package feedback

import (
	"moneh/models"
	"moneh/utils"

	"github.com/google/uuid"
)

// Feedback Interface
type FeedbackService interface {
	CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error
	GetAllFeedback(pagination utils.Pagination) ([]models.AllFeedbackData, int, error)
	HardDeleteFeedbackByID(ID uuid.UUID) error
}

// Feedback Struct
type feedbackService struct {
	feedbackRepo FeedbackRepository
}

// Feedback Constructor
func NewFeedbackService(feedbackRepo FeedbackRepository) FeedbackService {
	return &feedbackService{
		feedbackRepo: feedbackRepo,
	}
}

func (r *feedbackService) GetAllFeedback(pagination utils.Pagination) ([]models.AllFeedbackData, int, error) {
	return r.feedbackRepo.FindAllFeedback(pagination)
}

func (r *feedbackService) CreateFeedback(feedback *models.Feedback, userID uuid.UUID) error {
	return r.feedbackRepo.CreateFeedback(feedback, userID)
}

func (r *feedbackService) HardDeleteFeedbackByID(ID uuid.UUID) error {
	return r.feedbackRepo.HardDeleteFeedbackByID(ID)
}
