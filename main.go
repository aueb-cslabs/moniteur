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

	plugin.Initialize()
	rest.Initialize()

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(types.NewContext(c, plugin, calendar))
		}
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))

	e.GET("/api/schedule/all", rest.ScheduleAll)
	e.GET("/api/schedule/:room", rest.ScheduleRoom)
	e.GET("/api/schedule/:room/now", rest.ScheduleRoomNow)
	e.GET("/api/calendarInfo", rest.CalendarInfo)
	e.POST("/api/announcement", rest.CreateAnnouncement)
	e.DELETE("/api/announcement", rest.DeleteAnnouncement)
	e.PUT("/api/announcement", rest.UpdateAnnouncement)
	e.GET("/api/announcement", rest.Announcement)
	e.POST("/api/announcement/:room", rest.CreateRoomAnn)
	e.GET("/api/announcement/:room", rest.GetRoomAnn)
	e.DELETE("/api/announcement/:room", rest.DeleteRoomAnn)
	e.PUT("/api/announcement/:room", rest.UpdateRoomAnn)
	e.POST("/api/comment", rest.CreateComment)
	e.DELETE("/api/comment", rest.DeleteComment)
	e.PUT("/api/comment", rest.UpdateComment)
	e.GET("/api/comment", rest.Comment)

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
