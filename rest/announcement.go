package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
)

var message *types.Announcement

func AnnouncementsGroup(g *echo.Group) {
	g.POST("", createAnnouncement)
	g.DELETE("", deleteAnnouncement)
	g.PUT("", updateAnnouncement)
	g.GET("", announcement)
	g.POST("/:room", createRoomAnn)
	g.GET("/:room", getRoomAnn)
	g.DELETE("/:room", deleteRoomAnn)
	g.PUT("/:room", updateRoomAnn)
}

func createAnnouncement(e echo.Context) error {
	_ = e.Bind(&message)

	return e.NoContent(http.StatusOK)
}

func deleteAnnouncement(e echo.Context) error {
	message = nil

	return e.NoContent(http.StatusOK)
}

func announcement(e echo.Context) error {

	return e.JSON(http.StatusOK, announcement)
}

func updateAnnouncement(e echo.Context) error {
	_ = e.Bind(&message)

	return e.NoContent(http.StatusOK)
}
