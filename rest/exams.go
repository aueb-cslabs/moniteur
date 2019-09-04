package rest

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func ExamsGroup(g *echo.Group) {
	g.GET("/all", examsScheduleAll)
	g.GET("/:room", examsScheduleRoomToday)
	g.GET("/:room/now", examsScheduleRoomTodayNow)
}

func examsScheduleAll(ec echo.Context) error {
	c := ec.(*types.Context)
	schedule, err := c.Plugin().ExamsSchedule()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, schedule)
}

func examsScheduleRoomToday(ec echo.Context) error {
	c := ec.(*types.Context)

	schedule, err, room := c.Plugin().ExamsScheduleRoom(c.Param("room"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	t := time.Now()
	_, month, dayNum := t.Date()

	var applicableSlots []*types.ScheduleSlot
	for _, slot := range schedule.Slots {
		if slot.Room == room && slot.MonthNum == int(month) && slot.DayNum == dayNum {
			applicableSlots = append(applicableSlots, slot)
		}
	}
	schedule.Slots = applicableSlots

	return c.JSON(http.StatusOK, schedule)
}

func examsScheduleRoomTodayNow(ec echo.Context) error {
	c := ec.(*types.Context)

	schedule, err, room := c.Plugin().ExamsScheduleRoom(c.Param("room"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	t := time.Now()
	_, month, dayNum := t.Date()

	day, sec := determineNow()

	returnSchedule := &types.ScheduleNow{}
	for _, slot := range schedule.Slots {
		if slot.Room == room && slot.MonthNum == int(month) && slot.DayNum == dayNum {
			if slot.Day == day && slot.Start < sec && slot.End > sec {
				returnSchedule.Now = slot
			}
			if slot.Day == day && slot.Start >= sec {
				returnSchedule.Next = append(returnSchedule.Next, slot)
			}
		}
	}

	return c.JSON(http.StatusOK, returnSchedule)
}
