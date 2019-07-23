package types

import "time"

// Schedule The schedule object containing all the schedule slots.
type Schedule struct {
	Slots []*ScheduleSlot `json:"slots"`
}

/*
ScheduleSlot Contains information about a schedule slot, like times, room
and subject.
*/
type ScheduleSlot struct {
	Room  string
	Start time.Time
	End   time.Time
	Title string
	Host  string
}
