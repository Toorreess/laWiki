package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, ic IEntryController) *echo.Echo {
	api := e.Group("/api")

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))

	api.GET("/entry/:id", func(c echo.Context) error {
		return ic.Get(c)
	})

	api.POST("/entry", func(c echo.Context) error {
		return ic.Create(c)
	})

	api.PUT("/entry/:id", func(c echo.Context) error {
		obj := make(map[string]interface{})
		c.Bind(&obj)
		return ic.Update(c, obj)
	})

	api.DELETE("/entry/:id", func(c echo.Context) error {
		return ic.Delete(c)
	})

	api.GET("/entry", func(c echo.Context) error {
		return ic.List(c)
	})

	api.POST("/entry/:id/set-latest/:version_id", func(c echo.Context) error {
		return ic.SetLatest(c)
	})

	return e
}
