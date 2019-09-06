package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
)

// Map that connects rooms and announcements
var announcements map[string]*types.Announcement

// Initialize Method
func Initialize() {
	announcements = make(map[string]*types.Announcement)
}

// createRoomAnn Method that accepts POSTs a room announcement
func createRoomAnn(e echo.Context) error {
	post := &types.AnnouncementPost{}
	_ = e.Bind(post)

	ann := &types.Announcement{}
	ann.End = types.ConvertDateToUnix(post.End)
	ann.Msg = post.Msg

	announcements[e.Param("room")] = ann

	return e.NoContent(http.StatusOK)
}

// updateRoomAnn Method that accepts PUTs a room announcement
func updateRoomAnn(e echo.Context) error {
	post := &types.AnnouncementPost{}
	_ = e.Bind(post)

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
