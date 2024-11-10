package internal

import (
	"net/http"
	"strconv"

	"github.com/Toorreess/laWiki/entry-service/internal/model"
	"github.com/labstack/echo/v4"
)

type Context echo.Context

type IEntryController interface {
	Create(c Context) error
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

func (e *entryController) Create(c Context) error {
	var em *model.Entry
	if err := c.Bind(&em); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	em, err := e.EntryInteractor.Create(em)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	return c.JSON(http.StatusOK, em)
}

func (e *entryController) Get(c Context) error {
	id := c.Param("id")
	em, err := e.EntryInteractor.Get(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, em)
}

func (e *entryController) Update(c Context, body map[string]interface{}) error {
	var em *model.Entry

	em, err := e.EntryInteractor.Update(c.Param("id"), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, em)
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

	q := make(map[string]string)
	for k, v := range query {
		if k != "limit" && k != "offset" && k != "orderBy" && k != "order" {
			q[k] = v[0]
		}
	}

	limitStr := query.Get("limit")
	if limitStr == "" {
		limitStr = "20"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"status": "Limit must be a number"})
	}

	offsetStr := query.Get("offset")
	if offsetStr == "" {
		offsetStr = "0"
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"status": "Offset must be a number"})
	}

	list, err := e.EntryInteractor.List(q, limit, offset, query.Get("orderBy"), query.Get("order"))
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
