package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Toorreess/laWiki/api-gateway-service/internal/models"
	"github.com/labstack/echo/v4"
)

var ENTRY_SERVICE_HOST = os.Getenv("ENTRY_SERVICE_HOST")

func CreateEntry(c Context) error {
	var payload *models.Entry

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	jsonBytes, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, ENTRY_SERVICE_HOST, bytes.NewReader(jsonBytes))
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

func GetEntry(c Context) error {
	id := c.Param("id")

	req, err := http.NewRequest(http.MethodGet, ENTRY_SERVICE_HOST+"/"+id, nil)
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

func UpdateEntry(c Context) error {
	id := c.Param("id")
	var payload *models.Entry

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	jsonBytes, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPut, ENTRY_SERVICE_HOST+"/"+id, bytes.NewReader(jsonBytes))
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

func DeleteEntry(c Context) error {
	id := c.Param("id")

	req, err := http.NewRequest(http.MethodDelete, ENTRY_SERVICE_HOST+"/"+id, nil)
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

func ListEntry(c Context) error {
	req, err := http.NewRequest(http.MethodGet, ENTRY_SERVICE_HOST, nil)
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
