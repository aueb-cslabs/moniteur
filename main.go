// +build linux darwin

package main

import (
	"github.com/aueb-cslabs/moniteur/rest"
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/url"
	"strings"
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
	calendar, err := types.LoadCalendar("calendar.yml")
	if err != nil {
		log.Panic(err)
	}
	rooms, err := types.LoadMapping("mapping.yml")

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(types.NewContext(c, plugin, rooms, calendar))
		}
	})

	e.GET("/api/schedule/all", rest.ScheduleAll)
	e.GET("/api/schedule/:room", rest.ScheduleRoom)
	e.GET("/api/schedule/:room/now", rest.ScheduleRoomNow)
	e.GET("/api/info", rest.CalendarInfo)

	// Should go in effect only in development mode.
	// In production this should just serve the files.
	u, err := url.Parse("http://localhost:3000")
	if err != nil {
		e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{{URL: u}}
	proxyConfig := middleware.ProxyConfig{
		Skipper: func(c echo.Context) bool {
			return strings.HasPrefix(c.Path(), "/api")
		},
		Balancer: middleware.NewRandomBalancer(targets),
	}
	e.Use(middleware.ProxyWithConfig(proxyConfig))
	// End block

	e.Logger.Fatal(e.Start(":1323"))

}
