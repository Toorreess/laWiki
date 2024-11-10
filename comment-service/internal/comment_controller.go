package internal

import (
	"net/http"
	"strconv"

	"github.com/Toorreess/laWiki/comment-service/internal/model"
	"github.com/labstack/echo/v4"
)

type Context echo.Context

type ICommentController interface {
	Create(c Context) error
	Get(c Context) error
	Update(c Context, body map[string]interface{}) error
	Delete(c Context) error

	List(c Context) error
}

type commentController struct {
	CommentInteractor ICommentInteractor
}

func NewCommentController(ci ICommentInteractor) ICommentController {
	return &commentController{ci}
}

func (cc *commentController) Create(c Context) error {
	var cm *model.Comment

	if err := c.Bind(&cm); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	cm, err := cc.CommentInteractor.Create(cm)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	return c.JSON(http.StatusOK, cm)
}

func (cc *commentController) Get(c Context) error {
	id := c.Param("id")

	comment, err := cc.CommentInteractor.Get(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, comment)
}

func (cc *commentController) Update(c Context, body map[string]interface{}) error {
	var cm *model.Comment

	cm, err := cc.CommentInteractor.Update(c.Param("id"), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, cm)
}

func (cc *commentController) Delete(c Context) error {
	id := c.Param("id")

	if err := cc.CommentInteractor.Delete(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.NoContent(http.StatusNoContent)
}

func (cc *commentController) List(c Context) error {
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

	list, err := cc.CommentInteractor.List(q, limit, offset, query.Get("orderBy"), query.Get("order"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]string{"status": "Not found"})
	}

	return c.JSON(http.StatusOK, list)
}
