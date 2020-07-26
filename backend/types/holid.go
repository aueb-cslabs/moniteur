package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Holiday struct that defines a holiday
type Holiday struct {
	Date    string `json:"date"`
	Name    string `json:"localName"`
	IntName string `json:"name"`
}

func (h Holiday) Validate() error {
	return validation.ValidateStruct(&h,
		validation.Field(&h.Name, validation.Required),
		validation.Field(&h.IntName, validation.Required),
		validation.Field(&h.Date, validation.Required, validation.Date("2006-01-02")),
	)
}
