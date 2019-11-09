package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/utils"
	"github.com/labstack/echo"
	"net/http"
)

func Register(e echo.Context) error {
	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	if authorizedUsers[username] == "Y" {
		return e.NoContent(http.StatusAccepted)
	}
	authorizedUsers[username] = "Y"
	err := utils.UpdateUsers(authorizedUsers)
	if err != nil {
		return e.NoContent(http.StatusInternalServerError)
	}
	return e.NoContent(http.StatusOK)
}

func Unregister(e echo.Context) error {
	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	delete(authorizedUsers, username)
	err := utils.UpdateUsers(authorizedUsers)
	if err != nil {
		return e.NoContent(http.StatusInternalServerError)
	}
	return e.NoContent(http.StatusOK)
}
