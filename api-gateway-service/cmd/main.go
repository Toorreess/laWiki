package main

import (
	"github.com/Toorreess/laWiki/api-gateway-service/internal/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e = api.NewRouter(e)
	e.Logger.Fatal(e.Start(":1232"))
}
