package types

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jinzhu/gorm"
)

// User basic struct of a user
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"string;not_null"`
	Password string `json:"password" gorm:"-"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Username, validation.Required),
	)
}
