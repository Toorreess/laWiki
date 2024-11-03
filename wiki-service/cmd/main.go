package main

import (
	"log"

	wiki "github.com/Toorreess/laWiki/wiki-service/internal"
	"github.com/Toorreess/laWiki/wiki-service/internal/config"
	"github.com/Toorreess/laWiki/wiki-service/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.ReadConfig()

	db, err := database.NewDBClient(cfg.Database.DBType, cfg.Database.User, cfg.Database.Password, cfg.Database.Addr, cfg.Database.DBName)
	if err != nil {
		log.Fatalf("error initializing DB Client: %v\n", err)
	}

	defer db.Close()

	wikiController := wiki.NewWikiController(wiki.NewWikiInteractor(wiki.NewWikiRepository(db), wiki.NewWikiPresenter()))
	e := echo.New()
	e = wiki.NewRouter(e, wikiController)

	e.Logger.Fatal(e.Start(":1232"))
}
