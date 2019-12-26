package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/aueb-cslabs/moniteur/backend/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

//TODO REMOVE EXCESS CODE AND CLEAN IT UP

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
	com.End, err = types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	com.Msg = post.Msg
	redisClient.Set("Comment", com.Msg, 0)
	redisClient.ExpireAt("Comment", time.Unix(com.End, 0))
	redisClient.Set("Comment_dt", com.End, 0)

	writeComment()

	return e.NoContent(http.StatusOK)
}

// deleteComment Method that accepts DELETEs a general comment
func deleteComment(e echo.Context) error {
	com = nil

	writeComment()
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
		com = &types.Announcement{End: end, Msg: res}
		return e.JSON(http.StatusOK, com)
	} else {
		return e.JSON(http.StatusOK, nil)
	}
}

// updateComment Method that accepts PUTs a general comment
func updateComment(e echo.Context) error {
	post := types.AnnouncementPost{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	com = &types.Announcement{}
	com.End, err = types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	com.Msg = post.Msg
	redisClient.Set("Comment", com.Msg, 0)
	redisClient.ExpireAt("Comment", time.Unix(com.End, 0))
	redisClient.Set("Comment_dt", com.End, 0)

	writeComment()

	return e.NoContent(http.StatusOK)
}

// checkCommentExpiration checks every hour if the current comment has expired
func checkCommentExpiration() {
	for {
		if com != nil {
			now := time.Now().Unix()
			if com.End <= now {
				com = nil
			}
		}
		time.Sleep(time.Hour)
	}
}

// writeComment writes comment to existing.yml
func writeComment() {
	_ = utils.WriteComment(com)
}
