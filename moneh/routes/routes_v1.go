package routes

import (
	flwhandlers "moneh/modules/flows/http_handlers"
	pckhandlers "moneh/modules/pockets/http_handlers"
	stshandlers "moneh/modules/stats/http_handlers"
	syshandlers "moneh/modules/systems/http_handlers"
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

	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)
	e.DELETE("api/v1/dct/destroy/:id", syshandlers.HardDelDictionaryById)
	e.POST("api/v1/dct", syshandlers.PostDictionary)

	// Tags
	e.GET("api/v1/tag/:ord", syshandlers.GetAllTags)
	e.DELETE("api/v1/tag/destroy/:id", syshandlers.HardDelTagById)

	// Flows
	e.GET("api/v1/flows/:ord", flwhandlers.GetAllFlow)
	e.GET("api/v1/flows/summary/:type", flwhandlers.GetSummaryByType)
	e.POST("api/v1/flows", flwhandlers.PostFlow)
	e.DELETE("api/v1/flows/destroy/:id", flwhandlers.HardDelFlowById)
	e.DELETE("api/v1/flows/by/:id", flwhandlers.SoftDelFlowById)

	// Pockets
	e.GET("api/v1/pockets/headers/:ord", pckhandlers.GetAllPocketHeaders)
	e.POST("api/v1/pockets", pckhandlers.PostPocket)
	e.DELETE("api/v1/pockets/destroy/:id", pckhandlers.HardDelPocketById)

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
	e.GET("api/v1/stats/dcttype/:ord", stshandlers.GetTotalDctByType)
	e.GET("api/v1/stats/wishlisttype/:ord", stshandlers.GetTotalWishlistType)
	e.GET("api/v1/stats/wishlistpriority/:ord", stshandlers.GetTotalWishlistPriority)
	e.GET("api/v1/stats/wishlistisachieved/:ord", stshandlers.GetTotalWishlistIsAchieved)

	// =============== Private routes ===============

	return e
}
