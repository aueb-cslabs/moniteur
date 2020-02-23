package types

import "github.com/jinzhu/gorm"

type RoomSchedule struct {
	gorm.Model
	Room  string `gorm:"text;not_null"`
	Date  string `json:"date" gorm:"text;not_null"`
	Open  string `json:"open" gorm:"text;not_null"`
	Close string `json:"close" gorm:"text;not_null"`
}

type RoomScheduleArray struct {
	Rs []RoomSchedule `json:"rows"`
}
