package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func ScheduleAll(ec echo.Context) error {
	c := ec.(*types.Context)
	schedule, err, _ := c.Plugin().Schedule("")
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, schedule)
}

func ScheduleRoom(ec echo.Context) error {
	c := ec.(*types.Context)

	schedule, err, room := c.Plugin().Schedule(c.Param("room"))
	if err != nil {
		return err
	}

	var applicableSlots []*types.ScheduleSlot
	for _, slot := range schedule.Slots {
		if slot.Room == room {
			applicableSlots = append(applicableSlots, slot)
		}
	}
	schedule.Slots = applicableSlots

	return c.JSON(http.StatusOK, schedule)
}

func ScheduleRoomNow(ec echo.Context) error {
	c := ec.(*types.Context)

	schedule, err, room := c.Plugin().Schedule(c.Param("room"))
	if err != nil {
		return err
	}

	returnSchedule := &types.ScheduleNow{}

	day, sec := determineNow()

	for _, slot := range schedule.Slots {
		if slot.Room == room && slot.Day == day && slot.Start < sec && slot.End > sec {
			returnSchedule.Now = slot
		}
		if slot.Room == room && slot.Day == day && slot.Start >= sec {
			returnSchedule.Next = append(returnSchedule.Next, slot)
		}
	}

	return c.JSON(http.StatusOK, returnSchedule)
}

func determineNow() (int, int64) {
	now := time.Now()
	year, month, day := now.Date()
	sod := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return int(now.Weekday()), int64(now.Sub(sod).Seconds())
}
