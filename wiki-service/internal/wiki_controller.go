package internal

import (
	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
)

type Context echo.Context

type IWikiController interface {
	Create(c Context, authClient *auth.Client) error
	Read(c Context) error
	Update(c Context, body map[string]interface{}, authClient *auth.Client) error
	Delete(c Context, authClient *auth.Client) error

	List(c Context) error
}

type wikiController struct {
	WikiInteractor IWikiInteractor
}

func NewWikiController(wi IWikiInteractor) IWikiController {
	return &wikiController{wi}
}

func (w *wikiController) Create(c Context, authClient *auth.Client) error {
	panic("unimplemented")
}

func (w *wikiController) Read(c Context) error {
	panic("unimplemented")
}

func (w *wikiController) Update(c Context, body map[string]interface{}, authClient *auth.Client) error {
	panic("unimplemented")
}

func (w *wikiController) Delete(c Context, authClient *auth.Client) error {
	panic("unimplemented")
}

func (w *wikiController) List(c Context) error {
	panic("unimplemented")
}
