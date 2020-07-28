package authentication

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/aueb-cslabs/moniteur/backend/types"
)

func Register(e echo.Context) error {
	ctx := e.(*types.Context)
	db := ctx.DB

	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}

	user := &types.User{Username: username}

	if !db.Where("username = ?", username).Find(user).RecordNotFound() {
		return e.NoContent(http.StatusAccepted)
	}
	if err := db.Create(user).Error; err != nil {
		return e.NoContent(http.StatusBadRequest)
	}
	return e.NoContent(http.StatusOK)
}

func Unregister(e echo.Context) error {
	ctx := e.(*types.Context)
	db := ctx.DB

	username := e.Param("id")

	if username == "" {
		return e.NoContent(http.StatusBadRequest)
	}

	user := &types.User{Username: username}

	if err := db.Unscoped().Delete(user).Error; err != nil {
		return e.NoContent(http.StatusBadRequest)
	}
	return e.NoContent(http.StatusOK)
}
