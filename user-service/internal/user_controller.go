package internal

import (
	"net/http"
	"strconv"

	"github.com/Toorreess/laWiki/user-service/internal/model"
	"github.com/labstack/echo/v4"
)

type Context echo.Context

type IUserController interface {
	Create(c Context) error
	Get(c Context) error
	Update(c Context, body map[string]interface{}) error
	Delete(c Context) error
	List(c Context) error
}

type userController struct {
	UserInteractor IUserInteractor
}

func NewUserController(ui IUserInteractor) IUserController {
	return &userController{ui}
}

func (uc *userController) Create(c Context) error {
	var um *model.User

	if err := c.Bind(&um); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	um, err := uc.UserInteractor.Create(um)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	return c.JSON(http.StatusCreated, um)
}

func (uc *userController) Get(c Context) error {
	var um *model.User

	um, err := uc.UserInteractor.Get(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, um)
}

func (uc *userController) Update(c Context, body map[string]interface{}) error {
	var um *model.User

	if _, ok := body["id"]; ok {
		delete(body, "id")
	}

	um, err := uc.UserInteractor.Update(c.Param("id"), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, um)
}

func (uc *userController) Delete(c Context) error {
	if err := uc.UserInteractor.Delete(c.Param("id")); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.NoContent(http.StatusNoContent)
}

func (uc *userController) List(c Context) error {
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

	list, err := uc.UserInteractor.List(filteredQueryParams, limit, offset, orderBy, order)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, list)
}
