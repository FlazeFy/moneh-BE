package routes

import (
	flwhandlers "moneh/modules/flows/http_handlers"
	stshandlers "moneh/modules/stats/http_handlers"
	syshandlers "moneh/modules/systems/http_handlers"
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

	// Flows
	e.GET("api/v1/flows", flwhandlers.GetAllFlow)
	e.POST("api/v1/flows", flwhandlers.PostFlow)
	e.DELETE("api/v1/flows/destroy/:id", flwhandlers.HardDelFlowById)
	e.DELETE("api/v1/flows/by/:id", flwhandlers.SoftDelFlowById)

	// Stats
	e.GET("api/v1/stats/flowtype/:ord", stshandlers.GetTotalFlowByType)
	e.GET("api/v1/stats/flowcat/:ord", stshandlers.GetTotalFlowByCat)
	e.GET("api/v1/stats/dcttype/:ord", stshandlers.GetTotalDctByType)

	// =============== Private routes ===============

	return e
}
