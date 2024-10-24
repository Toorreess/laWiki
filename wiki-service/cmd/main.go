package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	wiki "github.com/Toorreess/laWiki/wiki-service/internal"
	"github.com/Toorreess/laWiki/wiki-service/pkg/config"
	"github.com/Toorreess/laWiki/wiki-service/pkg/database"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	cfg := config.ReadConfig()

	db, err := database.NewDBClient(cfg.Database.DBType, cfg.Database.User, cfg.Database.Password, cfg.Database.Addr, cfg.Database.DBName)
	if err != nil {
		log.Fatalf("error initializing DB Client: %v\n", err)
	}
	defer db.Client.(database.DBClient).Close()
	fbConfig := firebase.Config{
		ProjectID: cfg.Database.Addr,
	}

	app, err := firebase.NewApp(ctx, &fbConfig)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	fbClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error initializing Firebase Auth client: %v\n", err)
	}
	_ = fbClient

	e := echo.New()
	wikiController := wiki.NewWikiController(wiki.NewWikiInteractor(wiki.NewWikiRepository(db), wiki.NewWikiPresenter()))
	ro := wiki.NewRouter(e, wikiController, fbClient)

	_ = ro
}
