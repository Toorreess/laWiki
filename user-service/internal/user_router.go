package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, uc IUserController) *echo.Echo {
	api := e.Group("/api")

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))

	api.GET("/users/:id", func(c echo.Context) error {
		return uc.Get(c)
	})

	api.POST("/users", func(c echo.Context) error {
		return uc.Create(c)
	})

	api.PUT("/users/:id", func(c echo.Context) error {
		obj := make(map[string]interface{})
		c.Bind(&obj)
		return uc.Update(c, obj)
	})

	api.DELETE("/users/:id", func(c echo.Context) error {
		return uc.Delete(c)
	})

	api.GET("/users", func(c echo.Context) error {
		return uc.List(c)
	})

	return e
}
