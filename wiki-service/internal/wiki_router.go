package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, wc IWikiController) *echo.Echo {
	api := e.Group("/api")

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))

	api.GET("/wiki/:id", func(c echo.Context) error {
		return wc.Get(c)
	})

	api.POST("/wiki", func(c echo.Context) error {
		return wc.Create(c)
	})

	api.PUT("/wiki/:id", func(c echo.Context) error {
		obj := make(map[string]interface{})
		c.Bind(&obj)
		return wc.Update(c, obj)
	})

	api.DELETE("/wiki/:id", func(c echo.Context) error {
		return wc.Delete(c)
	})

	api.GET("/wiki", func(c echo.Context) error {
		return wc.List(c)
	})

	return e
}
