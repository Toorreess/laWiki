package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, cc ICommentController) *echo.Echo {
	api := e.Group("/api")

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))

	api.GET("/comment/:id", func(c echo.Context) error {
		return cc.Get(c)
	})

	api.POST("/comment", func(c echo.Context) error {
		return cc.Create(c)
	})

	api.PUT("/comment/:id", func(c echo.Context) error {
		obj := make(map[string]interface{})
		c.Bind(&obj)
		return cc.Update(c, obj)
	})

	api.DELETE("/comment/:id", func(c echo.Context) error {
		return cc.Delete(c)
	})

	api.GET("/comment", func(c echo.Context) error {
		return cc.List(c)
	})

	return e
}
