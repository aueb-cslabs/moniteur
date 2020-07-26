package databases

import (
	"os"

	"github.com/labstack/gommon/log"

	"github.com/aueb-cslabs/moniteur/backend/types"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitializePostgres() (db *gorm.DB) {
	overrides, err := gorm.Open("postgres", os.Getenv("POSTGRES_URI"))

	if err != nil {
		log.Panic(err)
	}

	overrides.AutoMigrate(&types.DBScheduleSlot{}, &types.User{})
	return overrides
}

func Terminate(db *gorm.DB) {
	_ = db.Close()
}
