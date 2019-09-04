package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
)

var message *types.Announcement

// AnnouncementsGroup Defines the api paths for all the announcements
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

// createAnnouncement Method that accepts POSTs a general announcement
func createAnnouncement(e echo.Context) error {
	_ = e.Bind(&message)

	return e.NoContent(http.StatusOK)
}

// deleteAnnouncement Method that accepts DELETEs a general announcement
func deleteAnnouncement(e echo.Context) error {
	message = nil

	return e.NoContent(http.StatusOK)
}

// announcement Method that accepts GETs a general announcement
func announcement(e echo.Context) error {

	return e.JSON(http.StatusOK, message)
}

// updateAnnouncement Method that accepts PUTs a general announcement
func updateAnnouncement(e echo.Context) error {
	_ = e.Bind(&message)

	return e.NoContent(http.StatusOK)
}
