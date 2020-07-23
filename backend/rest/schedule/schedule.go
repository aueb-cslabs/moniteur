package schedule

import (
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/aueb-cslabs/moniteur/backend/types"
)

// ScheduleGroup Defines the api paths for the normal schedule
func ScheduleGroup(g *echo.Group) {
	g.GET("/all", scheduleAll)
	g.GET("/:room", scheduleRoom)
	g.GET("/:room/now", scheduleRoomNow)
}

// scheduleAll Method that returns the normal schedule
func scheduleAll(ec echo.Context) error {
	c := ec.(*types.Context)
	schedule, err := c.Plugin().Schedule()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, schedule)
}

// scheduleRoom Method that returns the normal schedule for that specific room
func scheduleRoom(ec echo.Context) error {
	c := ec.(*types.Context)

	schedule, err, room := c.Plugin().ScheduleRoom(c.Param("room"))
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

// scheduleRoomNow Method that returns the normal schedule for that specific room and this current time
func scheduleRoomNow(ec echo.Context) error {
	c := ec.(*types.Context)

	schedule, err, room := c.Plugin().ScheduleRoom(c.Param("room"))
	if err != nil {
		return err
	}

	returnSchedule := &types.ScheduleNow{}

	day, _ := determineNow()
	unix := time.Now().Unix()

	for _, slot := range schedule.Slots {
		if slot.Room == room && slot.Day == day && slot.Start < unix && slot.End > unix {
			returnSchedule.Now = slot
		}
		if slot.Room == room && slot.Day == day && slot.Start >= unix {
			returnSchedule.Next = append(returnSchedule.Next, slot)
		}

		sort.Slice(returnSchedule.Next, func(i, j int) bool {
			return returnSchedule.Next[i].Start < returnSchedule.Next[j].Start
		})
	}

	db := c.DB
	var overrides []*types.DBScheduleSlot
	db.Where("room = ?", room).Find(&overrides)
	now := time.Now().Unix()

	for _, slot := range overrides {
		if slot.Start < now && slot.End > now {
			returnSchedule.Now = &slot.ScheduleSlot
		}
	}

	return c.JSON(http.StatusOK, returnSchedule)
}

// determineNow Method that determines the current time as a Unix timestamp and the current weekday
func determineNow() (int, int64) {
	now := time.Now()
	year, month, day := now.Date()
	sod := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return int(now.Weekday()), int64(now.Sub(sod).Seconds())
}

// Room returns the room name based on the mapping
func Room(e echo.Context) error {
	c := e.(*types.Context)
	_, _, room := c.Plugin().ScheduleRoom(c.Param("id"))

	return c.JSON(http.StatusOK, room)
}

func Rooms(e echo.Context) error {
	c := e.(*types.Context)
	rooms := c.Plugin().GetRooms()
	sort.Strings(rooms)
	return c.JSON(http.StatusOK, rooms)
}
