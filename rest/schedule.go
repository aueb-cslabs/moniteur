package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func ScheduleAll(ec echo.Context) error {
	c := ec.(*types.Context)
	schedule, err := c.Plugin().Schedule()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, schedule)
}

func ScheduleRoom(ec echo.Context) error {
	c := ec.(*types.Context)
	schedule, err := c.Plugin().Schedule()
	if err != nil {
		return err
	}

	var applicableSlots []*types.ScheduleSlot
	for _, slot := range schedule.Slots {
		if slot.Room == c.Param("room") {
			applicableSlots = append(applicableSlots, slot)
		}
	}
	schedule.Slots = applicableSlots

	return c.JSON(http.StatusOK, schedule)
}

func ScheduleRoomNow(ec echo.Context) error {
	c := ec.(*types.Context)
	schedule, err := c.Plugin().Schedule()
	if err != nil {
		return err
	}

	day, sec := determineNow()

	var applicableSlots []*types.ScheduleSlot
	for _, slot := range schedule.Slots {
		if slot.Room == c.Param("room") && slot.Day == day && slot.Start < sec && slot.End > sec {
			applicableSlots = append(applicableSlots, slot)
		}
	}
	schedule.Slots = applicableSlots

	return c.JSON(http.StatusOK, schedule)
}

func determineNow() (int, int64) {
	now := time.Now()
	year, month, day := now.Date()
	sod := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return int(now.Weekday()), int64(now.Sub(sod).Seconds())
}