package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

var com *types.Announcement

// CommentGroup Defines the api paths for all the comments
func CommentGroup(g *echo.Group) {
	g.POST("", Validate(createComment))
	g.DELETE("", Validate(deleteComment))
	g.PUT("", Validate(updateComment))
	g.GET("", comment)
}

// createComment Method that accepts POSTs a general comment
func createComment(e echo.Context) error {
	post := types.AnnouncementPost{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	com = &types.Announcement{}
	com.End = types.ConvertDateToUnix(post.End)
	com.Msg = post.Msg

	return e.NoContent(http.StatusOK)
}

// deleteComment Method that accepts DELETEs a general comment
func deleteComment(e echo.Context) error {
	com = nil

	return e.NoContent(http.StatusOK)
}

// comment Method that accepts GETs a general comment
func comment(e echo.Context) error {

	return e.JSON(http.StatusOK, com)
}

// updateComment Method that accepts PUTs a general comment
func updateComment(e echo.Context) error {
	post := types.AnnouncementPost{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	com = &types.Announcement{}
	com.End = types.ConvertDateToUnix(post.End)
	com.Msg = post.Msg

	return e.NoContent(http.StatusOK)
}

func checkCommentExpiration() {
	for {
		if com != nil {
			now := time.Now().Unix()
			if com.End >= now {
				com = nil
			}
		}
		time.Sleep(time.Hour)
	}
}
