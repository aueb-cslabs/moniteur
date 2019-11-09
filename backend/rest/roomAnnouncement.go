package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/aueb-cslabs/moniteur/backend/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// createRoomAnn Method that accepts POSTs a room announcement
func createRoomAnn(e echo.Context) error {
	c := e.(*types.Context)

	post := &types.AnnouncementPost{}
	err := e.Bind(post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	ann := &types.Announcement{}
	ann.End, err = types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	ann.Msg = post.Msg

	//Because of the room names that are in Greek we need to convert it to English
	room := c.Plugin().ConvertNameInverse(e.Param("room"))

	announcements[room] = ann

	writeRoomAnnouncements()

	return e.NoContent(http.StatusOK)
}

// updateRoomAnn Method that accepts PUTs a room announcement
func updateRoomAnn(e echo.Context) error {
	c := e.(*types.Context)

	post := &types.AnnouncementPost{}
	err := e.Bind(post)

	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	ann := &types.Announcement{}
	ann.End, err = types.ConvertDateToUnix(post.End)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}
	ann.Msg = post.Msg

	//Because of the room names that are in Greek we need to convert it to English
	room := c.Plugin().ConvertNameInverse(e.Param("room"))

	announcements[room] = ann

	writeRoomAnnouncements()

	return e.NoContent(http.StatusOK)
}

// deleteRoomAnn Method that accepts DELETEs a room announcement
func deleteRoomAnn(e echo.Context) error {
	c := e.(*types.Context)
	room := c.Plugin().ConvertNameInverse(e.Param("room"))
	delete(announcements, room)

	writeRoomAnnouncements()

	return e.NoContent(http.StatusOK)
}

// getRoomAnn Method that accepts GETs a room announcement
func getRoomAnn(e echo.Context) error {
	c := e.(*types.Context)
	room := c.Plugin().ConvertNameInverse(e.Param("room"))

	return e.JSON(http.StatusOK, announcements[room])
}

// checkRoomAnnouncementsExpiration checks every hour if any room announcement has expired
func checkRoomAnnouncementsExpiration() {
	for {
		for i, k := range announcements {
			now := time.Now().Unix()
			if now >= k.End {
				delete(announcements, i)
			}
		}
		time.Sleep(time.Hour)
	}
}

// writeRoomAnnouncements writes room announcement to existing.yml
func writeRoomAnnouncements() {
	_ = utils.WriteRoomAnnouncement(announcements)
}
