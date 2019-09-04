package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
)

var com *types.Announcement

func CommentGroup(g *echo.Group) {
	g.POST("", createComment)
	g.DELETE("", deleteComment)
	g.PUT("", updateComment)
	g.GET("", comment)
}

func createComment(e echo.Context) error {
	_ = e.Bind(&com)

	return e.NoContent(http.StatusOK)
}

func deleteComment(e echo.Context) error {
	com = nil

	return e.NoContent(http.StatusOK)
}

func comment(e echo.Context) error {

	return e.JSON(http.StatusOK, com)
}

func updateComment(e echo.Context) error {
	_ = e.Bind(&com)

	return e.NoContent(http.StatusOK)
}
