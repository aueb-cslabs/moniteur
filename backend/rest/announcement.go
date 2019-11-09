package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/aueb-cslabs/moniteur/backend/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

var message *types.Announcement

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
	post := types.AnnouncementPost{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	message = &types.Announcement{}
	message.End, err = types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	message.Msg = post.Msg

	writeAnnouncement()

	return e.NoContent(http.StatusOK)
}

// deleteAnnouncement Method that accepts DELETEs a general announcement
func deleteAnnouncement(e echo.Context) error {
	message = nil

	writeAnnouncement()

	return e.NoContent(http.StatusOK)
}

// announcement Method that accepts GETs a general announcement
func announcement(e echo.Context) error {

	return e.JSON(http.StatusOK, message)
}

// updateAnnouncement Method that accepts PUTs a general announcement
func updateAnnouncement(e echo.Context) error {
	post := types.AnnouncementPost{}
	err := e.Bind(&post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	message = &types.Announcement{}
	message.End, err = types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	message.Msg = post.Msg

	writeAnnouncement()

	return e.NoContent(http.StatusOK)
}

// checkAnnouncementExpiration checks every hour if the current announcement has expired
func checkAnnouncementExpiration() {
	for {
		if message != nil {
			now := time.Now().Unix()
			if message.End <= now {
				message = nil
				writeAnnouncement()
			}
		}
		time.Sleep(time.Hour)
	}
}

// writeAnnouncement writes announcement to existing.yml
func writeAnnouncement() {
	_ = utils.WriteAnnouncement(message)
}
