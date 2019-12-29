package postgres

import (
	"fmt"
	"github.com/aueb-cslabs/moniteur/backend/types"

	//"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Initialize(uri string) (db *gorm.DB) {
	overrides, err := gorm.Open("postgres", uri)
	if err != nil {
		fmt.Println(err)
	}

	overrides.AutoMigrate(&types.ScheduleSlot{})
	return overrides
}

func Terminate(db *gorm.DB) {
	_ = db.Close()
}
