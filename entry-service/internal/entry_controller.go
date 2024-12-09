package internal

import (
	"net/http"
	"strconv"

	"firebase.google.com/go/v4/storage"
	"github.com/Toorreess/laWiki/entry-service/internal/model"
	"github.com/labstack/echo/v4"
)

type Context echo.Context

type IEntryController interface {
	Create(c Context, storageClient *storage.Client) error
	Get(c Context) error
	Update(c Context, body map[string]interface{}) error
	Delete(c Context) error
	List(c Context) error

	SetLatest(c Context) error
}

type entryController struct {
	EntryInteractor IEntryInteractor
}

func NewEntryController(ei IEntryInteractor) IEntryController {
	return &entryController{ei}
}

func (e *entryController) Create(c Context, storageClient *storage.Client) error {
	var em *model.Entry
	if err := c.Bind(&em); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	em, err := e.EntryInteractor.Create(em)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	return c.JSON(http.StatusCreated, em)
}

func (e *entryController) Get(c Context) error {
	var em *model.Entry

	em, err := e.EntryInteractor.Get(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, em)
}

func (e *entryController) Update(c Context, body map[string]interface{}) error {
	return c.JSON(http.StatusOK, nil)
}

func (e *entryController) Delete(c Context) error {
	id := c.Param("id")

	if err := e.EntryInteractor.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.NoContent(http.StatusNoContent)
}

func (e *entryController) List(c Context) error {
	query := c.QueryParams()

	var limitStr, offsetStr, orderBy, order string

	filteredQueryParams := make(map[string]string)
	for k, v := range query {
		switch k {
		case "limit":
			limitStr = v[0]
		case "offset":
			offsetStr = v[0]
		case "orderBy":
			orderBy = v[0]
		case "order":
			order = v[0]
		default:
			filteredQueryParams[k] = v[0]
		}
	}

	if limitStr == "" {
		limitStr = "20"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"status": "Limit must be a number"})
	}

	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"status": "Offset must be a number"})
	}

	list, err := e.EntryInteractor.List(filteredQueryParams, limit, offset, orderBy, order)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, list)
}

func (e *entryController) SetLatest(c Context) error {
	entry_id := c.Param("id")
	version_id := c.Param("version_id")

	err := e.EntryInteractor.SetLatest(entry_id, version_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.NoContent(http.StatusNoContent)
}
