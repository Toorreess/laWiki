package main

import (
	"log"

	"github.com/Toorreess/laWiki/comment-service/config"
	comment "github.com/Toorreess/laWiki/comment-service/internal"
	"github.com/Toorreess/laWiki/comment-service/internal/database"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.ReadConfig()

	db, err := database.NewDBClient(cfg.Database.DBType, cfg.ProjectID)
	if err != nil {
		log.Fatalf("error initializing DB Client: %v\n", err)
	}

	defer db.Close()

	commentController := comment.NewCommentController(comment.NewcommentInteractor(comment.NewCommentRepository(db), comment.NewCommentPresenter()))
	e := echo.New()
	e = comment.NewRouter(e, commentController)

	e.Logger.Fatal(e.Start(":1232"))
}
