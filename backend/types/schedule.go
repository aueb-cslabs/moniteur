package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
)

// Schedule The schedule object containing all the schedule slots.
type Schedule struct {
	Slots []*ScheduleSlot `json:"slots"`
}

/*
ScheduleSlot Contains information about a schedule slot, like times, room
and subject.
*/
type ScheduleSlot struct {
	//ID int `gorm:"AUTO_INCREMENT;primary_key;not_null;unique_index"`
	Room string `json:"room" gorm:"text;not_null"`
	// Day is the days after the start of the week (0 = Sunday)
	Day int `json:"day" gorm:"-"`
	// Start is seconds after the start of day.
	Start int64 `json:"start" gorm:"numeric;not_null"`
	// End is seconds after the start of day.
	End      int64  `json:"end" gorm:"numeric;not_null"`
	Title    string `json:"title" gorm:"text;not_null"`
	Host     string `json:"host" gorm:"text;not_null"`
	Semester string `json:"semester" gorm:"-"`
	DayNum   int    `json:"-" gorm:"-"`
	MonthNum int    `json:"-" gorm:"-"`
	IsExam   bool   `json:"isExam" gorm:"boolean;not_null"`
}

type DBScheduleSlot struct {
	gorm.Model `json:"db_info"`
	ScheduleSlot
}

type ScheduleNow struct {
	Now  *ScheduleSlot   `json:"now"`
	Next []*ScheduleSlot `json:"next"`
}

func (ss ScheduleSlot) Validate() error {
	return validation.ValidateStruct(&ss,
		validation.Field(&ss.Title, validation.Required),
		validation.Field(&ss.Host, validation.Required),
		validation.Field(&ss.Start, validation.Required),
		validation.Field(&ss.End, validation.Required),
		validation.Field(&ss.Room, validation.Required),
		validation.Field(&ss.IsExam, validation.Required),
	)
}

func (dbss DBScheduleSlot) Validate() error {
	return dbss.ScheduleSlot.Validate()
}
