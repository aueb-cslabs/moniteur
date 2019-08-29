package rest

import (
	"fmt"
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
)

var comment *types.Announcement

func CreateComment(e echo.Context) error {
	_ = e.Bind(&comment)

	fmt.Println(comment)

	return e.NoContent(http.StatusOK)
}

func DeleteComment(e echo.Context) error {
	comment = nil

	return e.NoContent(http.StatusOK)
}

func Comment(e echo.Context) error {

	return e.JSON(http.StatusOK, comment)
}

func UpdateComment(e echo.Context) error {
	_ = e.Bind(&comment)

	return e.NoContent(http.StatusOK)
}
