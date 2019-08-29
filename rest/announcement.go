package rest

import (
	"fmt"
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
)

var announcement *types.Announcement

func CreateAnnouncement(e echo.Context) error {
	_ = e.Bind(&announcement)

	fmt.Println(announcement)

	return e.NoContent(http.StatusOK)
}

func DeleteAnnouncement(e echo.Context) error {
	announcement = nil

	return e.NoContent(http.StatusOK)
}

func Announcement(e echo.Context) error {

	return e.JSON(http.StatusOK, announcement)
}

func UpdateAnnouncement(e echo.Context) error {
	_ = e.Bind(&announcement)

	return e.NoContent(http.StatusOK)
}
