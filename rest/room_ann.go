package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
)

var announcements map[string]*types.Announcement

func Initialize() {
	announcements = make(map[string]*types.Announcement)
}

func createRoomAnn(e echo.Context) error {
	ann := &types.Announcement{}
	_ = e.Bind(ann)

	announcements[e.Param("room")] = ann

	return e.NoContent(http.StatusOK)
}

func updateRoomAnn(e echo.Context) error {
	ann := &types.Announcement{}
	_ = e.Bind(ann)

	announcements[e.Param("room")] = ann

	return e.NoContent(http.StatusOK)
}

func deleteRoomAnn(e echo.Context) error {
	delete(announcements, e.Param("room"))

	return e.NoContent(http.StatusOK)
}

func getRoomAnn(e echo.Context) error {
	return e.JSON(http.StatusOK, announcements[e.Param("room")])
}
