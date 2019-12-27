package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// AnnouncementsGroup Defines the api paths for all the announcements
func AnnouncementsGroup(g *echo.Group) {
	g.POST("", Validate(createAnnouncement))
	g.DELETE("", Validate(deleteAnnouncement))
	g.PUT("", Validate(updateAnnouncement))
	g.GET("", announcement)
	g.POST("/:room", Validate(createRoomAnn))
	g.GET("/:room", getRoomAnn)
	g.DELETE("/:room", Validate(deleteRoomAnn))
	g.PUT("/:room", Validate(updateRoomAnn))
}

// createAnnouncement Method that accepts POSTs a general announcement
func createAnnouncement(e echo.Context) error {
	post := types.Announcement{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	redisClient.Set("Announcement", post.Msg, 0)
	redisClient.ExpireAt("Announcement", time.Unix(post.End, 0))
	redisClient.Set("Announcement_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}

// deleteAnnouncement Method that accepts DELETEs a general announcement
func deleteAnnouncement(e echo.Context) error {
	redisClient.Del("Announcement")
	redisClient.Del("Announcement_dt")

	return e.NoContent(http.StatusOK)
}

// announcement Method that accepts GETs a general announcement
func announcement(e echo.Context) error {
	exists := redisClient.Exists("Announcement").Val()
	if exists == 1 {
		res := redisClient.Get("Announcement").Val()
		end, _ := redisClient.Get("Announcement_dt").Int64()
		message := &types.Announcement{End: end, Msg: res}
		return e.JSON(http.StatusOK, message)
	} else {
		return e.JSON(http.StatusOK, nil)
	}
}

// updateAnnouncement Method that accepts PUTs a general announcement
func updateAnnouncement(e echo.Context) error {
	post := types.Announcement{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	redisClient.Set("Announcement", post.Msg, 0)
	redisClient.ExpireAt("Announcement", time.Unix(post.End, 0))
	redisClient.Set("Announcement_dt", post.End, 0)

	return e.NoContent(http.StatusOK)
}
