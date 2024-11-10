package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Toorreess/laWiki/api-gateway-service/config"
	"github.com/Toorreess/laWiki/api-gateway-service/internal/models"
	"github.com/labstack/echo/v4"
)

var commentPort = config.ReadConfig().Server.CommentPort
var COMMENT_SERVICE_HOST = fmt.Sprintf("http://comment-service%s/api/comment", commentPort)

func CreateComment(c Context) error {
	var payload *models.Comment

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	jsonBytes, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, COMMENT_SERVICE_HOST, bytes.NewReader(jsonBytes))
	if err != nil {
		return echo.ErrInternalServerError
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Response().Writer, resp.Body)

	return nil
}

func GetComment(c Context) error {
	id := c.Param("id")

	req, err := http.NewRequest(http.MethodGet, COMMENT_SERVICE_HOST+"/"+id, nil)
	if err != nil {
		return echo.ErrInternalServerError
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.ErrServiceUnavailable
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)

	io.Copy(c.Response().Writer, resp.Body)
	return nil
}

func UpdateComment(c Context) error {
	id := c.Param("id")
	var payload *models.Comment

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	jsonBytes, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPut, COMMENT_SERVICE_HOST+"/"+id, bytes.NewReader(jsonBytes))
	if err != nil {
		return echo.ErrInternalServerError
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.ErrInternalServerError
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Response().Writer, resp.Body)

	return nil
}

func DeleteComment(c Context) error {
	id := c.Param("id")

	req, err := http.NewRequest(http.MethodDelete, COMMENT_SERVICE_HOST+"/"+id, nil)
	if err != nil {
		return echo.ErrInternalServerError
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.ErrServiceUnavailable
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)

	return nil
}

func ListComment(c Context) error {
	req, err := http.NewRequest(http.MethodGet, COMMENT_SERVICE_HOST, nil)
	if err != nil {
		return echo.ErrInternalServerError
	}

	req.URL.Query().Set("limit", c.QueryParam("limit"))
	req.URL.Query().Set("offset", c.QueryParam("offset"))
	req.URL.Query().Set("orderBy", c.QueryParam("orderBy"))
	req.URL.Query().Set("order", c.QueryParam("order"))

	if c.QueryParam("name") != "" {
		req.URL.Query().Set("name", c.QueryParam("name"))
	}

	if c.QueryParam("author") != "" {
		req.URL.Query().Set("author", c.QueryParam("author"))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.ErrServiceUnavailable
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Response().Writer, resp.Body)

	return nil
}
