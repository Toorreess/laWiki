package main

import (
	"github.com/Toorreess/laWiki/api-gateway-service/config"
	"github.com/Toorreess/laWiki/api-gateway-service/internal/api"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.ReadConfig()

	e := echo.New()
	e = api.NewRouter(e)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
