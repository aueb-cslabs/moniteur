package announcements

import (
	"github.com/aueb-cslabs/moniteur/backend/rest/authentication"
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// CommentGroup Defines the api paths for all the comments
func CommentGroup(g *echo.Group) {
	g.POST("", authentication.Validate(createComment))
	g.DELETE("", authentication.Validate(deleteComment))
	g.PUT("", authentication.Validate(updateComment))
	g.GET("", comment)
}

// createComment Method that accepts POSTs a general comment
func createComment(e echo.Context) error {
	ctx := e.(*types.Context)

	post := types.Announcement{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	ctx.RedisClient.Set("Comment", post.Msg, 0)
	ctx.RedisClient.ExpireAt("Comment", time.Unix(post.End, 0))
	ctx.RedisClient.Set("Comment_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}

// deleteComment Method that accepts DELETEs a general comment
func deleteComment(e echo.Context) error {
	ctx := e.(*types.Context)

	ctx.RedisClient.Del("Comment")
	ctx.RedisClient.Del("Comment_dt")

	return e.NoContent(http.StatusOK)
}

// comment Method that accepts GETs a general comment
func comment(e echo.Context) error {
	ctx := e.(*types.Context)

	exists := ctx.RedisClient.Exists("Comment").Val()
	if exists == 1 {
		res := ctx.RedisClient.Get("Comment").Val()
		end, _ := ctx.RedisClient.Get("Comment_dt").Int64()
		com := &types.Announcement{End: end, Msg: res}
		return e.JSON(http.StatusOK, com)
	} else {
		return e.JSON(http.StatusOK, nil)
	}
}

// updateComment Method that accepts PUTs a general comment
func updateComment(e echo.Context) error {
	ctx := e.(*types.Context)

	post := types.Announcement{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	ctx.RedisClient.Set("Comment", post.Msg, 0)
	ctx.RedisClient.ExpireAt("Comment", time.Unix(post.End, 0))
	ctx.RedisClient.Set("Comment_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}
