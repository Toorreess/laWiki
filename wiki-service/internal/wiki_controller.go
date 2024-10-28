package internal

import (
	"github.com/labstack/echo/v4"
)

type Context echo.Context

type IWikiController interface {
	Create(c Context) error
	Get(c Context) error
	Update(c Context, body map[string]interface{}) error
	Delete(c Context) error

	List(c Context) error
}

type wikiController struct {
	WikiInteractor IWikiInteractor
}

func NewWikiController(wi IWikiInteractor) IWikiController {
	return &wikiController{wi}
}

func (w *wikiController) Create(c Context) error {
	panic("unimplemented")
}

func (w *wikiController) Get(c Context) error {
	panic("unimplemented")
}

func (w *wikiController) Update(c Context, body map[string]interface{}) error {
	panic("unimplemented")
}

func (w *wikiController) Delete(c Context) error {
	panic("unimplemented")
}

func (w *wikiController) List(c Context) error {
	panic("unimplemented")
}
