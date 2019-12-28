package rest

import (
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/labstack/echo"
	"net/http"
)

func OverrideGroup(g *echo.Group) {
	g.POST("", Validate(createOverride))
	g.DELETE("", Validate(deleteOverride))
}

func createOverride(e echo.Context) error {
	db := e.(*types.Context).DB

	override := &types.ScheduleSlot{}
	err := e.Bind(override)

	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	db.Create(override)
	return e.NoContent(http.StatusOK)
}

func deleteOverride(e echo.Context) error {
	db := e.(*types.Context).DB

	override := &types.ScheduleSlot{}
	err := e.Bind(override)

	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	db.Delete(types.ScheduleSlot{}, "id = ?", override.Model.ID)
	return e.NoContent(http.StatusOK)
}
