package rest

import (
	"github.com/labstack/echo"
	"net/http"
)

func Register(e echo.Context) error {
	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	if authUsers.Get(username).Val() == "Y" {
		return e.NoContent(http.StatusAccepted)
	}
	_, err := authUsers.Set(username, "Y", 0).Result()
	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}
	return e.NoContent(http.StatusOK)
}

func Unregister(e echo.Context) error {
	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	res, _ := authUsers.Del(username).Result()
	if res != 1 {
		return e.NoContent(http.StatusBadRequest)
	}
	return e.NoContent(http.StatusOK)
}
