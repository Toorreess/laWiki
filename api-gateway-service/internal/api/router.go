package api

import (
	"github.com/Toorreess/laWiki/api-gateway-service/internal/handlers"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	api := e.Group("/api")
	v1 := api.Group("/v1")

	/* Wiki microservice endpoints */
	v1.POST("/wikis", func(c echo.Context) error {
		return handlers.CreateWiki(c)
	})
	v1.GET("/wikis/:id", func(c echo.Context) error {
		return handlers.GetWiki(c)
	})
	v1.PUT("/wikis/:id", func(c echo.Context) error {
		return handlers.UpdateWiki(c)
	})
	v1.DELETE("/wikis/:id", func(c echo.Context) error {
		return handlers.DeleteWiki(c)
	})
	v1.GET("/wikis", func(c echo.Context) error {
		return handlers.ListWiki(c)
	})

	/* Entry microservice endpoints */
	v1.POST("/wikis/:wiki_id/entries", func(c echo.Context) error {
		return handlers.CreateEntry(c)
	})
	v1.GET("/entries/:id", func(c echo.Context) error {
		return handlers.GetEntry(c)
	})
	v1.PUT("/entries/:id", func(c echo.Context) error {
		return handlers.UpdateEntry(c)
	})
	v1.DELETE("/entries/:id", func(c echo.Context) error {
		return handlers.DeleteEntry(c)
	})
	v1.GET("/entries", func(c echo.Context) error {
		return handlers.ListEntry(c)
	})

	/* Comment microservice endpoints */
	v1.POST("/entries/:entry_id/comments", func(c echo.Context) error {
		return handlers.CreateComment(c)
	})
	v1.GET("/comments/:id", func(c echo.Context) error {
		return handlers.GetComment(c)
	})
	v1.PUT("/comments/:id", func(c echo.Context) error {
		return handlers.UpdateComment(c)
	})
	v1.DELETE("/comments/:id", func(c echo.Context) error {
		return handlers.DeleteComment(c)
	})
	v1.GET("/comments", func(c echo.Context) error {
		return handlers.ListComment(c)
	})

	/* User microservice endpoints */
	v1.POST("/users", func(c echo.Context) error {
		return handlers.CreateUser(c)
	})
	v1.GET("/users/:id", func(c echo.Context) error {
		return handlers.GetUser(c)
	})
	v1.PUT("/users/:id", func(c echo.Context) error {
		return handlers.UpdateUser(c)
	})
	v1.DELETE("/users/:id", func(c echo.Context) error {
		return handlers.DeleteUser(c)
	})
	v1.GET("/users", func(c echo.Context) error {
		return handlers.ListUser(c)
	})

	v1.GET("/users/:id/notifications", func(c echo.Context) error {
		return handlers.GetNotifications(c)
	})

	v1.POST("/users/:id/notifications", func(c echo.Context) error {
		return handlers.AddNotification(c)
	})

	v1.PUT("/users/:user_id/notifications/:notification_id", func(c echo.Context) error {
		return handlers.ReadNotification(c)
	})

	return e
}
