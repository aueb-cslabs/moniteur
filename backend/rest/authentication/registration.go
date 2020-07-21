package authentication

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/aueb-cslabs/moniteur/backend/types"
)

func Register(e echo.Context) error {
	ctx := e.(*types.Context)

	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	if ctx.AuthUsers.Get(username).Val() == "Y" {
		return e.NoContent(http.StatusAccepted)
	}
	_, err := ctx.AuthUsers.Set(username, "Y", 0).Result()
	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}
	return e.NoContent(http.StatusOK)
}

func Unregister(e echo.Context) error {
	ctx := e.(*types.Context)

	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}
	res, _ := ctx.AuthUsers.Del(username).Result()
	if res != 1 {
		return e.NoContent(http.StatusBadRequest)
	}
	return e.NoContent(http.StatusOK)
}
