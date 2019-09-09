package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
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
	ann.End = types.ConvertDateToUnix(post.End)
	ann.Msg = post.Msg

	announcements[e.Param("room")] = ann

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
	ann.End = types.ConvertDateToUnix(post.End)
	ann.Msg = post.Msg

	announcements[e.Param("room")] = ann

	return e.NoContent(http.StatusOK)
}

// deleteRoomAnn Method that accepts DELETEs a room announcement
func deleteRoomAnn(e echo.Context) error {
	delete(announcements, e.Param("room"))

	return e.NoContent(http.StatusOK)
}

// getRoomAnn Method that accepts GETs a room announcement
func getRoomAnn(e echo.Context) error {
	return e.JSON(http.StatusOK, announcements[e.Param("room")])
}

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
