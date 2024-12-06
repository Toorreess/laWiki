package main

import (
	"log"

	"github.com/Toorreess/laWiki/user-service/config"
	user "github.com/Toorreess/laWiki/user-service/internal"
	"github.com/Toorreess/laWiki/user-service/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.ReadConfig()

	db, err := database.NewDBClient(cfg.Database.DBType, cfg.ProjectID)
	if err != nil {
		log.Fatalf("error initializing DB Client: %v\n", err)
	}

	defer db.Close()

	userController := user.NewUserController(user.NewUserInteractor(user.NewUserRepository(db), user.NewUserPresenter()))
	e := echo.New()
	e = user.NewRouter(e, userController)

	e.Logger.Fatal(e.Start(cfg.Server.Port))
}
