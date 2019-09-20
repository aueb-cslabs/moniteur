package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/aueb-cslabs/moniteur/utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// createRoomAnn Method that accepts POSTs a room announcement
func createRoomAnn(e echo.Context) error {
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

	announcements[e.Param("room")] = ann

	writeRoomAnnouncements()

	return e.NoContent(http.StatusOK)
}

// updateRoomAnn Method that accepts PUTs a room announcement
func updateRoomAnn(e echo.Context) error {
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

	announcements[e.Param("room")] = ann

	writeRoomAnnouncements()

	return e.NoContent(http.StatusOK)
}

// deleteRoomAnn Method that accepts DELETEs a room announcement
func deleteRoomAnn(e echo.Context) error {
	delete(announcements, e.Param("room"))

	writeRoomAnnouncements()

	return e.NoContent(http.StatusOK)
}

// getRoomAnn Method that accepts GETs a room announcement
func getRoomAnn(e echo.Context) error {
	return e.JSON(http.StatusOK, announcements[e.Param("room")])
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
