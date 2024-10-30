package internal

import (
	"github.com/labstack/echo/v4"
)

type Context echo.Context

type IEntryController interface {
	Create(c Context) error
	Get(c Context) error
	Update(c Context, body map[string]interface{}) error
	Delete(c Context) error

	List(c Context) error
}

type entryController struct {
	EntryInteractor IEntryInteractor
}

func NewEntryController(ei IEntryInteractor) IEntryController {
	return &entryController{ei}
}

func (e *entryController) Create(c Context) error {
	panic("unimplemented")
}

func (e *entryController) Get(c Context) error {
	panic("unimplemented")
}

func (e *entryController) Update(c Context, body map[string]interface{}) error {
	panic("unimplemented")
}

func (e *entryController) Delete(c Context) error {
	panic("unimplemented")
}

func (e *entryController) List(c Context) error {
	panic("unimplemented")
}
