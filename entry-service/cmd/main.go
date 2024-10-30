package main

import (
	"log"

	entry "github.com/Toorreess/laWiki/entry-service/internal"
	"github.com/Toorreess/laWiki/entry-service/internal/config"
	"github.com/Toorreess/laWiki/entry-service/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.ReadConfig()

	db, err := database.NewDBClient(cfg.Database.DBType, cfg.Database.User, cfg.Database.Password, cfg.Database.Addr, cfg.Database.DBName)
	if err != nil {
		log.Fatalf("error initializing DB Client: %v\n", err)
	}
	defer db.Client.(database.DBClient).Close()

	entryController := entry.NewEntryController(entry.NewEntryInteractor(entry.NewEntryRepository(db), entry.NewEntryPresenter()))
	e := echo.New()
	e = entry.NewRouter(e, entryController)

	e.Logger.Fatal(e.Start(":1232"))
}
