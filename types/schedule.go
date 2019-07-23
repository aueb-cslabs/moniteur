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
	Room  string
	// Day is the days after the start of the week (0 = Sunday)
	Day int
	// Start is seconds after the start of day.
	Start int64
	// End is seconds after the start of day.
	End   int64
	Title string
	Host  string
}
