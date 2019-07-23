// +build windows

package main

import (
	"github.com/aueb-cslabs/moniteur/aueb"
	"github.com/labstack/echo"
	"net/http"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		schedule, err := aueb.PLUGIN.Schedule()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, schedule)
	})
	e.Logger.Fatal(e.Start(":1323"))

}
