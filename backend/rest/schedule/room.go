package schedule

import (
	"github.com/aueb-cslabs/moniteur/backend/rest/authentication"
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/labstack/echo"
	"net/http"
)

func RoomSchedule(g *echo.Group) {
	g.GET("/:room", roomSchedule)
	g.POST("/:room", authentication.Validate(addRoomSchedule))
	g.DELETE("/:room", authentication.Validate(removeRoomSchedule))
}

func roomSchedule(e echo.Context) error {
	db := e.(*types.Context).DB

	var roomSchedule []*types.RoomSchedule

	db.Where("room = ?", e.Param("room")).Find(&roomSchedule)

	return e.JSON(http.StatusOK, roomSchedule)
}

func addRoomSchedule(e echo.Context) error {
	db := e.(*types.Context).DB

	rs := &types.RoomSchedule{}
	err := e.Bind(rs)
	rs.Room = e.Param("room")

	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	db.Create(rs)
	return e.NoContent(http.StatusOK)
}

func removeRoomSchedule(e echo.Context) error {
	db := e.(*types.Context).DB

	rs := &types.RoomScheduleArray{}
	err := e.Bind(rs)

	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	for _, sch := range rs.Rs {
		db.Delete(types.RoomSchedule{}, "id = ?", sch.Model.ID)
	}

	return e.NoContent(http.StatusOK)
}
