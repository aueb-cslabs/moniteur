package schedule

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/aueb-cslabs/moniteur/backend/plugin/aueb"
	"github.com/aueb-cslabs/moniteur/backend/types"
)

// ExamsGroup Defines the api paths for all the exams
func ExamsGroup(g *echo.Group) {
	g.GET("/all", examsScheduleAll)
	g.GET("/:room", examsScheduleRoomToday)
	g.GET("/:room/now", examsScheduleRoomTodayNow)
}

// examsScheduleAll Method that returns the exams schedule
func examsScheduleAll(ec echo.Context) error {
	c := ec.(*types.Context)
	schedule, err := aueb.ExamsSchedule()
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, schedule)
}

// examsScheduleRoomToday Method that returns the exams schedule for today and that specific room
func examsScheduleRoomToday(ec echo.Context) error {
	c := ec.(*types.Context)

	schedule, err, room := aueb.ExamsScheduleRoom(c.Param("room"))
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

// examsScheduleRoomTodayNow Method that returns the exams schedule for today, the current time and that specific room
func examsScheduleRoomTodayNow(ec echo.Context) error {
	c := ec.(*types.Context)

	schedule, err, room := aueb.ExamsScheduleRoom(c.Param("room"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	t := time.Now()
	unix := t.Unix()
	_, month, dayNum := t.Date()

	day, _ := determineNow()

	returnSchedule := &types.ScheduleNow{}
	for _, slot := range schedule.Slots {
		if slot.Room == room && slot.MonthNum == int(month) && slot.DayNum == dayNum {
			if slot.Day == day && slot.Start < unix && slot.End > unix {
				returnSchedule.Now = slot
			}
			if slot.Day == day && slot.Start >= unix {
				returnSchedule.Next = append(returnSchedule.Next, slot)
			}
		}
	}

	return c.JSON(http.StatusOK, returnSchedule)
}
