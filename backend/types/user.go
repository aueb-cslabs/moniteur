package types

import "github.com/jinzhu/gorm"

// User basic struct of a user
type User struct {
	gorm.Model
	Username string `json:"username" gorm:"type:text;PRIMARY_KEY"`
	Password string `json:"password" gorm:"-"`
}
