package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
)

func Register(e echo.Context) error {
	ctx := e.(*types.Context)
	username := e.Param("username")

	if username != "" {
		return e.NoContent(http.StatusBadRequest)
	}
	res, err := ctx.Plugin().RegisterAuthorizedUser(username)
	if err != nil || res != true {
		return e.JSON(http.StatusNoContent, err)
	}
	return e.NoContent(http.StatusOK)
}
