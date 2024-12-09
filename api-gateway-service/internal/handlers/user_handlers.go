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

var userPort = config.ReadConfig().Server.UserPort
var USER_SERVICE_HOST string = fmt.Sprintf("http://user-service%s/api/users", userPort)

func CreateUser(c Context) error {
	var payload *models.User

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	jsonBytes, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPost, USER_SERVICE_HOST, bytes.NewReader(jsonBytes))
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
func GetUser(c Context) error {
	id := c.Param("id")

	req, err := http.NewRequest(http.MethodGet, USER_SERVICE_HOST+"/"+id, nil)
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
func UpdateUser(c Context) error {
	id := c.Param("id")
	var payload *models.User

	if err := c.Bind(&payload); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	jsonBytes, err := json.Marshal(payload)
	req, err := http.NewRequest(http.MethodPut, USER_SERVICE_HOST+"/"+id, bytes.NewReader(jsonBytes))
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
func DeleteUser(c Context) error {
	id := c.Param("id")

	req, err := http.NewRequest(http.MethodDelete, USER_SERVICE_HOST+"/"+id, nil)
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
func ListUser(c Context) error {
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

	if c.QueryParam("email") != "" {
		req.URL.Query().Set("email", c.QueryParam("email"))
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

func GetNotifications(c Context) error {
	id := c.Param("id")

	req, err := http.NewRequest(http.MethodGet, USER_SERVICE_HOST+"/"+id, nil)
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

	var usr *models.User
	if err = json.NewDecoder(resp.Body).Decode(&usr); err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}

	json.NewEncoder(c.Response().Writer).Encode(usr.Notifications)
	return nil
}
func AddNotification(c Context) error {
	var notification *models.Notification
	if err := c.Bind(&notification); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, map[string]string{"status": "Not valid body"})
	}

	jsonBytes, err := json.Marshal(notification)
	req, err := http.NewRequest(http.MethodPost, USER_SERVICE_HOST+"/"+c.Param("id")+"/notifications", bytes.NewReader(jsonBytes))
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

func ReadNotification(c Context) error {
	userID := c.Param("user_id")
	notificationID := c.Param("notification_id")

	req, err := http.NewRequest(http.MethodPut, USER_SERVICE_HOST+"/"+userID+"/notifications/"+notificationID, nil)
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
