package internal

import (
	"net/http"
	"strconv"

	"github.com/Toorreess/laWiki/wiki-service/internal/model"
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
	var wm *model.Wiki

	if err := c.Bind(&wm); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	wm, err := w.WikiInteractor.Create(wm)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	return c.JSON(http.StatusCreated, wm)
}

func (w *wikiController) Get(c Context) error {
	var wm *model.Wiki

	wm, err := w.WikiInteractor.Get(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, wm)
}

func (w *wikiController) Update(c Context, body map[string]interface{}) error {
	var wm *model.Wiki

	wm, err := w.WikiInteractor.Update(c.Param("id"), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, wm)
}

func (w *wikiController) Delete(c Context) error {
	err := w.WikiInteractor.Delete(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.NoContent(http.StatusNoContent)
}

// TODO: No se muestra las entidades correctamente
func (w *wikiController) List(c Context) error {
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

	list, err := w.WikiInteractor.List(q, limit, offset, query.Get("orderBy"), query.Get("order"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, list)
}
