package main

import (
	"log"

	"github.com/Toorreess/laWiki/wiki-service/config"
	wiki "github.com/Toorreess/laWiki/wiki-service/internal"
	"github.com/Toorreess/laWiki/wiki-service/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.ReadConfig()

	db, err := database.NewDBClient(cfg.Database.DBType, cfg.ProjectID)
	if err != nil {
		log.Fatalf("error initializing DB Client: %v\n", err)
	}

	defer db.Close()

	wikiController := wiki.NewWikiController(wiki.NewWikiInteractor(wiki.NewWikiRepository(db), wiki.NewWikiPresenter()))
	e := echo.New()
	e = wiki.NewRouter(e, wikiController)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
