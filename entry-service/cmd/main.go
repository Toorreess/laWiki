package main

import (
	"log"

	"github.com/Toorreess/laWiki/entry-service/config"
	entry "github.com/Toorreess/laWiki/entry-service/internal"
	"github.com/Toorreess/laWiki/entry-service/internal/database"
	"github.com/labstack/echo/v4"

	firebase "firebase.google.com/go/v4"
)

func main() {
	cfg := config.ReadConfig()

	db, err := database.NewDBClient(cfg.Database.DBType, cfg.ProjectID)
	if err != nil {
		log.Fatalf("error initializing DB Client: %v\n", err)
	}

	defer db.Close()
	fbConfig := &firebase.Config{
		ProjectID:     cfg.ProjectID,
		StorageBucket: "lawiki-89989.appspot.com",
	}
	app, err := firebase.NewApp(db.Ctx, fbConfig)
	if err != nil {
		log.Fatalln(err)
	}

	storageClient, err := app.Storage(db.Ctx)
	if err != nil {
		log.Fatalln(err)
	}

	entryController := entry.NewEntryController(entry.NewEntryInteractor(entry.NewEntryRepository(db), entry.NewEntryPresenter()))

	e := echo.New()
	e = entry.NewRouter(e, entryController, storageClient)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
