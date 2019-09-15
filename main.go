// +build linux darwin

package main

import "C"
import (
	"github.com/aueb-cslabs/moniteur/rest"
	"github.com/aueb-cslabs/moniteur/types"
	"github.com/aueb-cslabs/moniteur/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/url"
	"strconv"
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

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	existing, _ := utils.Read()
	plugin.Initialize(config.ExamsLink)
	rest.Initialize(config.Secret, existing)

	api := e.Group("/api")

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(types.NewContext(c, plugin, calendar))
		}
	})

	//For CORS to work, please define the EXACT link that you will use!!!
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
	}))

	rest.CalendarGroup(api.Group("/calendarInfo"))
	rest.AnnouncementsGroup(api.Group("/announcement"))
	rest.CommentGroup(api.Group("/comment"))
	rest.ScheduleGroup(api.Group("/schedule"))
	rest.ExamsGroup(api.Group("/exams"))
	api.GET("/room/:id", rest.Room)
	api.POST("/authenticate", rest.Authenticate)

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

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.Port)))

}
