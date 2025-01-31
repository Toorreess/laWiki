package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Toorreess/laWiki/api-gateway-service/config"
	"github.com/Toorreess/laWiki/api-gateway-service/internal/models"
	"github.com/labstack/echo/v4"
)

type Context echo.Context

var wikiPort = config.ReadConfig().Server.WikiPort
var WIKI_SERVICE_HOST string = fmt.Sprintf("http://wiki-service%s/api/wikis", wikiPort)

func CreateWiki(c Context) error {
	var payload *models.Wiki

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	jsonBytes, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, WIKI_SERVICE_HOST, bytes.NewReader(jsonBytes))
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Response().Writer, resp.Body)

	return nil
}

func GetWiki(c Context) error {
	id := c.Param("id")

	req, err := http.NewRequest(http.MethodGet, WIKI_SERVICE_HOST+"/"+id, nil)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return echo.ErrServiceUnavailable
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)

	io.Copy(c.Response().Writer, resp.Body)
	return nil
}

func UpdateWiki(c Context) error {
	id := c.Param("id")
	var payload *models.Wiki

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	jsonBytes, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPut, WIKI_SERVICE_HOST+"/"+id, bytes.NewReader(jsonBytes))
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Response().Writer, resp.Body)

	return nil
}

func DeleteWiki(c Context) error {
	id := c.Param("id")

	req, err := http.NewRequest(http.MethodDelete, WIKI_SERVICE_HOST+"/"+id, nil)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return echo.ErrServiceUnavailable
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)

	return nil
}

func ListWiki(c Context) error {
	req, err := http.NewRequest(http.MethodGet, WIKI_SERVICE_HOST, nil)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return echo.ErrServiceUnavailable
	}
	defer resp.Body.Close()

	c.Response().Writer.Header().Set("Content-Type", "application/json")
	c.Response().Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Response().Writer, resp.Body)

	return nil
}
