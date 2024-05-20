package routes

import (
	middlewares "moneh/middlewares/jwt"
	authhandlers "moneh/modules/auth/http_handlers"
	flwhandlers "moneh/modules/flows/http_handlers"
	pckhandlers "moneh/modules/pockets/http_handlers"
	stshandlers "moneh/modules/stats/http_handlers"
	syshandlers "moneh/modules/systems/http_handlers"
	ushandlers "moneh/modules/users/http_handlers"
	wshhandlers "moneh/modules/wishlists/http_handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitV1() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Moneh")
	})

	// =============== Public routes ===============

	// Auth
	e.POST("api/v1/login", authhandlers.PostLoginUser)
	e.POST("api/v1/register", authhandlers.PostRegister)
	e.POST("api/v1/logout", authhandlers.SignOut)

	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)
	e.GET("api/v2/dct/:type", syshandlers.GetDictionaryByTypeFirebase)
	e.DELETE("api/v1/dct/destroy/:id", syshandlers.HardDelDictionaryById, middlewares.CustomJWTAuth)
	e.POST("api/v1/dct", syshandlers.PostDictionary, middlewares.CustomJWTAuth)

	// Tags
	e.GET("api/v1/tag/:ord", syshandlers.GetAllTags)
	e.DELETE("api/v1/tag/destroy/:id", syshandlers.HardDelTagById, middlewares.CustomJWTAuth)
	e.POST("api/v1/tag", syshandlers.PostTag, middlewares.CustomJWTAuth)

	e.GET("api/v2/tag/:ord", syshandlers.GetAllTagsFirebase)

	// Flows
	e.GET("api/v1/flows/:ord", flwhandlers.GetAllFlow)
	e.GET("api/v1/flows/month_item/:mon/:year/:type", flwhandlers.GetMonthlyFlowItem)
	e.GET("api/v1/flows/month_total/:mon/:year/:type", flwhandlers.GetMonthlyFlowTotal)
	e.GET("api/v1/flows/summary/:type", flwhandlers.GetSummaryByType)
	e.GET("api/v1/flows/dateammount/:type/:view", flwhandlers.GetTotalItemAmmountPerDateByType)
	e.POST("api/v1/flows", flwhandlers.PostFlow)
	e.DELETE("api/v1/flows/destroy/:id", flwhandlers.HardDelFlowById)
	e.DELETE("api/v1/flows/by/:id", flwhandlers.SoftDelFlowById)

	// Pockets
	e.GET("api/v1/pockets/headers/:ord", pckhandlers.GetAllPocketHeaders)
	e.POST("api/v1/pockets", pckhandlers.PostPocket)
	e.PUT("api/v1/pockets/by/:id", pckhandlers.UpdatePocket)
	e.DELETE("api/v1/pockets/destroy/:id", pckhandlers.HardDelPocketById)

	// Feedbacks
	e.POST("api/v1/feedbacks", syshandlers.PostFeedback)
	e.GET("api/v1/feedbacks/:ord_obj/:ord", syshandlers.GetAllFeedback)

	// Wishlists
	e.GET("api/v1/wishlists/headers/:ord", wshhandlers.GetAllWishlistHeaders)
	e.POST("api/v1/wishlists", wshhandlers.PostWishlist)
	e.DELETE("api/v1/wishlists/destroy/:id", wshhandlers.HardDelWishlistById)
	e.DELETE("api/v1/wishlists/by/:id", wshhandlers.SoftDelWishlistById)
	e.GET("api/v1/wishlists/summary", wshhandlers.GetSummary)

	// Stats
	e.GET("api/v1/stats/flowtype/:ord", stshandlers.GetTotalFlowByType)
	e.GET("api/v1/stats/flowcat/:ord", stshandlers.GetTotalFlowByCat)
	e.GET("api/v1/stats/pockettype/:ord", stshandlers.GetTotalPocketByType)
	e.GET("api/v1/stats/ammountflowtype/:ord", stshandlers.GetTotalAmmountFlowByType)
	e.GET("api/v1/stats/dcttype/:ord", stshandlers.GetTotalDctByType)
	e.GET("api/v1/stats/wishlisttype/:ord", stshandlers.GetTotalWishlistType)
	e.GET("api/v1/stats/wishlistpriority/:ord", stshandlers.GetTotalWishlistPriority)
	e.GET("api/v1/stats/wishlistisachieved/:ord", stshandlers.GetTotalWishlistIsAchieved)
	e.GET("api/v1/stats/dctmod/:table/:col", stshandlers.GetTotalDictionaryToModule)

	// Dashboard
	e.GET("api/v1/dashboard", stshandlers.GetDashboard)

	// =============== Private routes ===============
	// User
	e.GET("api/v1/user/my", ushandlers.GetMyProfile, middlewares.CustomJWTAuth)

	return e
}
