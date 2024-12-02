package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"firebase.google.com/go/v4/storage"
)

func NewRouter(e *echo.Echo, ic IEntryController, storageClient *storage.Client) *echo.Echo {
	api := e.Group("/api")

	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))

	api.GET("/entries/:id", func(c echo.Context) error {
		return ic.Get(c)
	})

	api.POST("/entries", func(c echo.Context) error {
		return ic.Create(c, storageClient)
	})

	api.PUT("/entries/:id", func(c echo.Context) error {
		obj := make(map[string]interface{})
		c.Bind(&obj)
		return ic.Update(c, obj)
	})

	api.DELETE("/entries/:id", func(c echo.Context) error {
		return ic.Delete(c)
	})

	api.GET("/entries", func(c echo.Context) error {
		return ic.List(c)
	})

	return e
}
