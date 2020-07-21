package databases

import (
	"github.com/labstack/gommon/log"

	"github.com/aueb-cslabs/moniteur/backend/types"

	"github.com/jinzhu/gorm"
)

func InitializePostgres(uri string) (db *gorm.DB) {
	overrides, err := gorm.Open("postgres", uri)
	if err != nil {
		log.Panic(err)
	}

	overrides.AutoMigrate(&types.DBScheduleSlot{})
	return overrides
}

func Terminate(db *gorm.DB) {
	_ = db.Close()
}
