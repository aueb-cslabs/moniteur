package announcements

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/aueb-cslabs/moniteur/backend/rest/authentication"
	"github.com/aueb-cslabs/moniteur/backend/types"
)

// AnnouncementsGroup Defines the api paths for all the announcements
func AnnouncementsGroup(g *echo.Group) {
	g.POST("", authentication.Validate(createAnnouncement))
	g.DELETE("", authentication.Validate(deleteAnnouncement))
	g.PUT("", authentication.Validate(updateAnnouncement))
	g.GET("", announcement)
	g.POST("/:room", authentication.Validate(createRoomAnn))
	g.GET("/:room", getRoomAnn)
	g.DELETE("/:room", authentication.Validate(deleteRoomAnn))
	g.PUT("/:room", authentication.Validate(updateRoomAnn))
}

// createAnnouncement Method that accepts POSTs a general announcement
func createAnnouncement(e echo.Context) error {
	ctx := e.(*types.Context)

	post := types.Announcement{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	ctx.RedisClient.Set("Announcement", post.Msg, 0)
	ctx.RedisClient.ExpireAt("Announcement", time.Unix(post.End, 0))
	ctx.RedisClient.Set("Announcement_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}

// deleteAnnouncement Method that accepts DELETEs a general announcement
func deleteAnnouncement(e echo.Context) error {
	ctx := e.(*types.Context)

	ctx.RedisClient.Del("Announcement")
	ctx.RedisClient.Del("Announcement_dt")

	return e.NoContent(http.StatusOK)
}

// announcement Method that accepts GETs a general announcement
func announcement(e echo.Context) error {
	ctx := e.(*types.Context)

	exists := ctx.RedisClient.Exists("Announcement").Val()
	if exists == 1 {
		res := ctx.RedisClient.Get("Announcement").Val()
		end, _ := ctx.RedisClient.Get("Announcement_dt").Int64()
		message := &types.Announcement{End: end, Msg: res}
		return e.JSON(http.StatusOK, message)
	} else {
		return e.JSON(http.StatusOK, nil)
	}
}

// updateAnnouncement Method that accepts PUTs a general announcement
func updateAnnouncement(e echo.Context) error {
	ctx := e.(*types.Context)

	post := types.Announcement{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	ctx.RedisClient.Set("Announcement", post.Msg, 0)
	ctx.RedisClient.ExpireAt("Announcement", time.Unix(post.End, 0))
	ctx.RedisClient.Set("Announcement_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}
