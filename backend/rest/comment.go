package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// CommentGroup Defines the api paths for all the comments
func CommentGroup(g *echo.Group) {
	g.POST("", Validate(createComment))
	g.DELETE("", Validate(deleteComment))
	g.PUT("", Validate(updateComment))
	g.GET("", comment)
}

// createComment Method that accepts POSTs a general comment
func createComment(e echo.Context) error {
	post := types.Announcement{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	redisClient.Set("Comment", post.Msg, 0)
	redisClient.ExpireAt("Comment", time.Unix(post.End, 0))
	redisClient.Set("Comment_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}

// deleteComment Method that accepts DELETEs a general comment
func deleteComment(e echo.Context) error {
	redisClient.Del("Comment")
	redisClient.Del("Comment_dt")

	return e.NoContent(http.StatusOK)
}

// comment Method that accepts GETs a general comment
func comment(e echo.Context) error {
	exists := redisClient.Exists("Comment").Val()
	if exists == 1 {
		res := redisClient.Get("Comment").Val()
		end, _ := redisClient.Get("Comment_dt").Int64()
		com := &types.Announcement{End: end, Msg: res}
		return e.JSON(http.StatusOK, com)
	} else {
		return e.JSON(http.StatusOK, nil)
	}
}

// updateComment Method that accepts PUTs a general comment
func updateComment(e echo.Context) error {
	post := types.Announcement{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	redisClient.Set("Comment", post.Msg, 0)
	redisClient.ExpireAt("Comment", time.Unix(post.End, 0))
	redisClient.Set("Comment_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}
