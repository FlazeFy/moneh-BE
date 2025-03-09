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
	e.POST("api/v1/logout", authhandlers.SignOut, middlewares.CustomJWTAuth)

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
	e.GET("api/v1/flows/:ord", flwhandlers.GetAllFlow, middlewares.CustomJWTAuth)
	e.GET("api/v2/flows", flwhandlers.GetAllFlowExport)
	e.GET("api/v1/flows/month_item/:mon/:year/:type", flwhandlers.GetMonthlyFlowItem, middlewares.CustomJWTAuth)
	e.GET("api/v1/flows/month_total/:mon/:year/:type", flwhandlers.GetMonthlyFlowTotal, middlewares.CustomJWTAuth)
	e.GET("api/v1/flows/summary/:type", flwhandlers.GetSummaryByType, middlewares.CustomJWTAuth)
	e.GET("api/v1/flows/dateammount/:type/:view", flwhandlers.GetTotalItemAmmountPerDateByType, middlewares.CustomJWTAuth)
	e.POST("api/v1/flows", flwhandlers.PostFlow, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/flows/destroy/:id", flwhandlers.HardDelFlowById, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/flows/by/:id", flwhandlers.SoftDelFlowById, middlewares.CustomJWTAuth)

	// Pockets
	e.GET("api/v1/pockets/headers/:ord", pckhandlers.GetAllPocketHeaders, middlewares.CustomJWTAuth)
	e.GET("api/v2/pockets", pckhandlers.GetAllPocketExport)
	e.POST("api/v1/pockets", pckhandlers.PostPocket, middlewares.CustomJWTAuth)
	e.PUT("api/v1/pockets/by/:id", pckhandlers.UpdatePocket, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/pockets/destroy/:id", pckhandlers.HardDelPocketById, middlewares.CustomJWTAuth)

	// Feedbacks
	e.POST("api/v1/feedbacks", syshandlers.PostFeedback)
	e.GET("api/v1/feedbacks/:ord_obj/:ord", syshandlers.GetAllFeedback)

	// Wishlists
	e.GET("api/v1/wishlists/headers/:ord", wshhandlers.GetAllWishlistHeaders, middlewares.CustomJWTAuth)
	e.POST("api/v1/wishlists", wshhandlers.PostWishlist, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/wishlists/destroy/:id", wshhandlers.HardDelWishlistById, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/wishlists/by/:id", wshhandlers.SoftDelWishlistById, middlewares.CustomJWTAuth)
	e.GET("api/v1/wishlists/summary", wshhandlers.GetSummary, middlewares.CustomJWTAuth)

	// Stats
	e.GET("api/v1/stats/summary/apps", stshandlers.GetSummaryApps)
	e.GET("api/v1/stats/flowtype/:ord", stshandlers.GetTotalFlowByType, middlewares.CustomJWTAuth)
	e.GET("api/v1/stats/flowcat/:ord", stshandlers.GetTotalFlowByCat, middlewares.CustomJWTAuth)
	e.GET("api/v1/stats/pockettype/:ord", stshandlers.GetTotalPocketByType, middlewares.CustomJWTAuth)
	e.GET("api/v1/stats/ammountflowtype/:ord", stshandlers.GetTotalAmmountFlowByType, middlewares.CustomJWTAuth)
	e.GET("api/v1/stats/wishlisttype/:ord", stshandlers.GetTotalWishlistType, middlewares.CustomJWTAuth)
	e.GET("api/v1/stats/wishlistpriority/:ord", stshandlers.GetTotalWishlistPriority, middlewares.CustomJWTAuth)
	e.GET("api/v1/stats/wishlistisachieved/:ord", stshandlers.GetTotalWishlistIsAchieved, middlewares.CustomJWTAuth)

	// Dashboard
	e.GET("api/v1/dashboard", stshandlers.GetDashboard, middlewares.CustomJWTAuth)

	// User
	e.GET("api/v1/user/my", ushandlers.GetMyProfile, middlewares.CustomJWTAuth)
	e.PUT("api/v1/user/telegram", ushandlers.UpdateTelegram, middlewares.CustomJWTAuth)

	return e
}
