package api

import (
	"github.com/Toorreess/laWiki/api-gateway-service/internal/handlers"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	/* Wiki microservice endpoints */
	v1.POST("/wiki", func(c echo.Context) error {
		return handlers.CreateWiki(c)
	})
	v1.GET("/wiki/:id", func(c echo.Context) error {
		return handlers.GetWiki(c)
	})
	v1.PUT("/wiki/:id", func(c echo.Context) error {
		return handlers.UpdateWiki(c)
	})
	v1.DELETE("/wiki/:id", func(c echo.Context) error {
		return handlers.DeleteWiki(c)
	})
	v1.GET("/wiki", func(c echo.Context) error {
		return handlers.ListWiki(c)
	})

	/* Entry microservice endpoints */
	v1.POST("/wiki/:wiki_id/entry", func(c echo.Context) error {
		return handlers.CreateEntry(c)
	})
	v1.GET("/entry/:id", func(c echo.Context) error {
		return handlers.GetEntry(c)
	})
	v1.PUT("/entry/:id", func(c echo.Context) error {
		return handlers.UpdateEntry(c)
	})
	v1.DELETE("/entry/:id", func(c echo.Context) error {
		return handlers.DeleteEntry(c)
	})
	v1.GET("/entry", func(c echo.Context) error {
		return handlers.ListEntry(c)
	})

	/* Comment microservice endpoints */
	v1.POST("/entry/:entry_id/comment", func(c echo.Context) error {
		return handlers.CreateComment(c)
	})
	v1.GET("/comment/:id", func(c echo.Context) error {
		return handlers.GetComment(c)
	})
	v1.PUT("/comment/:id", func(c echo.Context) error {
		return handlers.UpdateComment(c)
	})
	v1.DELETE("/comment/:id", func(c echo.Context) error {
		return handlers.DeleteComment(c)
	})
	v1.GET("/comment", func(c echo.Context) error {
		return handlers.ListComment(c)
	})

	return e
}
