package modules

import (
	"moneh/middlewares/business"
	"moneh/modules/admin"
	"moneh/modules/auth"
	"moneh/modules/dictionary"
	"moneh/modules/errors"
	"moneh/modules/feedback"
	"moneh/modules/flow"
	"moneh/modules/history"
	"moneh/modules/pocket"
	"moneh/modules/stats"
	"moneh/modules/user"
	"moneh/seeders"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetUpDependency(r *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	// Dependency Cache
	statsCache := stats.NewStatsCache(redisClient)

	// Dependency Repositories
	adminRepo := admin.NewAdminRepository(db)
	feedbackRepo := feedback.NewFeedbackRepository(db)
	historyRepo := history.NewHistoryRepository(db)
	userRepo := user.NewUserRepository(db)
	flowRepo := flow.NewFlowRepository(db)
	pocketRepo := pocket.NewPocketRepository(db)
	flowRelationRepo := flow.NewFlowRelationRepository(db)
	dictionaryRepo := dictionary.NewDictionaryRepository(db)
	errorRepo := errors.NewErrorRepository(db)
	statsRepo := stats.NewStatsRepository(db)

	// Dependency Services
	// adminService := services.NewAdminService(adminRepo)
	authService := auth.NewAuthService(userRepo, adminRepo, redisClient)
	feedbackService := feedback.NewFeedbackService(feedbackRepo)
	historyService := history.NewHistoryService(historyRepo)
	userService := user.NewUserService(userRepo)
	dictionaryService := dictionary.NewDictionaryService(dictionaryRepo)
	pocketService := pocket.NewPocketService(pocketRepo, historyRepo)
	flowService := flow.NewFlowService(flowRepo, historyRepo, flowRelationRepo)
	errorService := errors.NewErrorService(errorRepo)
	statsService := stats.NewStatsService(statsRepo, redisClient, statsCache)

	// Dependency Controller
	authController := auth.NewAuthController(authService)
	feedbackController := feedback.NewFeedbackController(feedbackService)
	historyController := history.NewHistoryController(historyService)
	userController := user.NewUserController(userService)
	dictionaryController := dictionary.NewDictionaryController(dictionaryService)
	pocketController := pocket.NewPocketController(pocketService, statsService)
	flowController := flow.NewFlowController(flowService, statsService)
	errorController := errors.NewErrorController(errorService)

	// Other Middleware
	currencyMiddleware := business.NewCurrencyMiddleware(userService)

	// Routes Endpoint
	auth.AuthRouter(r, redisClient, *authController)
	feedback.FeedbackRouter(r, *feedbackController, redisClient, db)
	history.HistoryRouter(r, *historyController, redisClient, db)
	user.UserRouter(r, *userController, redisClient)
	dictionary.DictionaryRouter(r, *dictionaryController, redisClient, db)
	pocket.PocketRouter(r, *pocketController, redisClient, currencyMiddleware, db)
	flow.FlowRouter(r, *flowController, redisClient, currencyMiddleware, db)
	errors.ErrorRouter(r, *errorController, redisClient, db)

	// Seeder & Factories
	seeders.SeedAdmins(adminRepo, 5)
	seeders.SeedUsers(userRepo, 20)
	seeders.SeedFeedbacks(feedbackRepo, userRepo, 15)
	seeders.SeedHistories(historyRepo, userRepo, 60)
	seeders.SeedFlows(flowRepo, userRepo, 300)
	seeders.SeedPockets(pocketRepo, userRepo, 60)
	seeders.SeedFlowRelations(flowRelationRepo, userRepo, flowRepo, pocketRepo, 3)
	seeders.SeedDictionaries(dictionaryRepo)
	seeders.SeedErrors(errorRepo, 25)
}
