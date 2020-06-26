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

	if err := user.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	newUser := &types.User{}
	if err := db.Where("username = ?", user.Username).First(&newUser).Error; err == nil {
		return e.NoContent(http.StatusBadRequest)
	}
	if err := db.Create(&types.User{Username: username}).Error; err != nil {
		return e.NoContent(http.StatusBadRequest)
	}
	return e.NoContent(http.StatusOK)
}

func Unregister(e echo.Context) error {
	ctx := e.(*types.Context)
	db := ctx.DB

	username := e.Param("id")
	user := &types.User{Username: username}

	if err := user.Validate(); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	if err := db.Where("username = ?", user.Username).Delete(types.User{}).Error; err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	return e.NoContent(http.StatusOK)
}
