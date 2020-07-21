package schedule

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/aueb-cslabs/moniteur/backend/rest/authentication"
	"github.com/aueb-cslabs/moniteur/backend/types"
)

func OverrideGroup(g *echo.Group) {
	g.POST("", authentication.Validate(createOverride))
	g.DELETE("", authentication.Validate(deleteOverride))
	g.GET("", authentication.Validate(overrides))
}

func createOverride(e echo.Context) error {
	db := e.(*types.Context).DB

	override := &types.DBScheduleSlot{}
	err := e.Bind(override)

	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	db.Create(override)
	return e.NoContent(http.StatusOK)
}

func deleteOverride(e echo.Context) error {
	db := e.(*types.Context).DB

	override := &types.DBScheduleSlot{}
	err := e.Bind(&override)

	if err != nil {
		return e.NoContent(http.StatusBadRequest)
	}

	db.Delete(types.DBScheduleSlot{}, "id = ?", override.Model.ID)
	return e.NoContent(http.StatusOK)
}

func overrides(e echo.Context) error {
	db := e.(*types.Context).DB

	var overrides []*types.DBScheduleSlot

	db.Find(&overrides)

	return e.JSON(http.StatusOK, overrides)
}
