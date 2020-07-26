package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Announcement contains the message to broadcast and the date that it expires
type Announcement struct {
	End int64  `json:"end" yaml:"end"`
	Msg string `json:"msg" yaml:"msg"`
}

func (a Announcement) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Msg, validation.Required),
		validation.Field(&a.End, validation.Required),
	)
}
