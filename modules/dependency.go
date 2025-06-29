package modules

import (
	"moneh/modules/admin"
	"moneh/modules/auth"
	"moneh/modules/feedback"
	"moneh/modules/history"
	"moneh/modules/user"
	"moneh/seeders"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetUpDependency(r *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	// Dependency Repositories
	adminRepo := admin.NewAdminRepository(db)
	feedbackRepo := feedback.NewFeedbackRepository(db)
	historyRepo := history.NewHistoryRepository(db)
	userRepo := user.NewUserRepository(db)

	// Dependency Services
	// adminService := services.NewAdminService(adminRepo)
	authService := auth.NewAuthService(userRepo, adminRepo, redisClient)
	feedbackService := feedback.NewFeedbackService(feedbackRepo)
	historyService := history.NewHistoryService(historyRepo)
	// userService := services.NewUserService(userRepo)

	// Dependency Controller
	authController := auth.NewAuthController(authService)
	feedbackController := feedback.NewFeedbackController(feedbackService)
	historyController := history.NewHistoryController(historyService)

	// Routes Endpoint
	auth.AuthRouter(r, redisClient, *authController)
	feedback.FeedbackRouter(r, *feedbackController)
	history.HistoryRouter(r, *historyController)

	// Seeder & Factories
	seeders.SeedAdmins(adminRepo, 5)
	seeders.SeedUsers(userRepo, 20)
	seeders.SeedFeedbacks(feedbackRepo, userRepo, 10)
	seeders.SeedHistories(historyRepo, userRepo, 60)
}
