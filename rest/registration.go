package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/aueb-cslabs/moniteur/utils"
	"github.com/labstack/echo"
	"net/http"
)

func Register(e echo.Context) error {
	ctx := e.(*types.Context)
	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	if ctx.AuthorizedUsers[username] == "Y" {
		return e.NoContent(http.StatusAccepted)
	}
	ctx.AuthorizedUsers[username] = "Y"
	err := utils.UpdateUsers(ctx.AuthorizedUsers)
	if err != nil {
		return e.NoContent(http.StatusInternalServerError)
	}
	return e.NoContent(http.StatusOK)
}

func Unregister(e echo.Context) error {
	ctx := e.(*types.Context)
	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	delete(ctx.AuthorizedUsers, username)
	err := utils.UpdateUsers(ctx.AuthorizedUsers)
	if err != nil {
		return e.NoContent(http.StatusInternalServerError)
	}
	return e.NoContent(http.StatusOK)
}
