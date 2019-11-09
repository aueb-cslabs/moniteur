package types

// Schedule The schedule object containing all the schedule slots.
type Schedule struct {
	Slots []*ScheduleSlot `json:"slots"`
}

/*
ScheduleSlot Contains information about a schedule slot, like times, room
and subject.
*/
type ScheduleSlot struct {
	Room string `json:"room"`
	// Day is the days after the start of the week (0 = Sunday)
	Day int `json:"day"`
	// Start is seconds after the start of day.
	Start int64 `json:"start"`
	// End is seconds after the start of day.
	End      int64  `json:"end"`
	Title    string `json:"title"`
	Host     string `json:"host"`
	Semester string `json:"semester"`
	DayNum   int    `json:"-"`
	MonthNum int    `json:"-"`
}

type ScheduleNow struct {
	Now  *ScheduleSlot   `json:"now"`
	Next []*ScheduleSlot `json:"next"`
}
