package main

import (
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

func main() {

	config, err := types.LoadConfiguration("config.yml")
	if err != nil {
		log.Panic(err)
	}
	plugin, err := types.LoadPlugin(config.Plugin)
	if err != nil {
		log.Panic(err)
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		schedule, err := plugin.Schedule()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, schedule)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
