package internal

import (
	"net/http"
	"strconv"

	"firebase.google.com/go/v4/storage"
	"github.com/labstack/echo/v4"
)

type Context echo.Context

type IEntryController interface {
	Create(c Context, storageClient *storage.Client) error
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

func (e *entryController) Create(c Context, storageClient *storage.Client) error {
	r := c.Request()

	r.ParseMultipartForm(10 << 20) // 10 MB limit

	entryData := map[string]string{
		"name":    r.FormValue("name"),
		"author":  r.FormValue("author"),
		"wiki_id": r.FormValue("wiki_id"),
	}

	file, _, err := r.FormFile("content")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"status": "File uploaded invalid"})
	}

	defer file.Close()

	resp, err := e.EntryInteractor.Create(entryData, file, storageClient)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"status": "File uploaded invalid"})
	}

	return c.JSON(http.StatusCreated, resp)
}

func (e *entryController) Get(c Context) error {
	panic("unimplemented")
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
