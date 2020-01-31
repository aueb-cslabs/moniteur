// +build linux darwin

package main

import "C"
import (
	"github.com/aueb-cslabs/moniteur/backend/databases"
	"github.com/aueb-cslabs/moniteur/backend/rest"
	"github.com/aueb-cslabs/moniteur/backend/rest/announcements"
	"github.com/aueb-cslabs/moniteur/backend/rest/authentication"
	"github.com/aueb-cslabs/moniteur/backend/rest/calendar"
	"github.com/aueb-cslabs/moniteur/backend/rest/schedule"
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tylerb/graceful"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {

	config, err := types.LoadConfiguration("config/config.yml")
	if err != nil {
		log.Panic(err)
	}
	plugin, err := types.LoadPlugin(config.Plugin)
	if err != nil {
		log.Panic(err)
	}
	calendarInfo, err := types.LoadCalendar("config/calendar.yml")
	if err != nil {
		log.Panic(err)
	}

	psql := databases.InitializePostgres(config.Postgres)
	redisClient, authUsers, tokens := databases.InitializeRedis(config.RedisConfig)

	plugin.Initialize(config.ExamsLink)
	rest.Initialize(calendarInfo)

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	api := e.Group("/api")

	calendar.CalendarGroup(api.Group("/calendarInfo"))
	announcements.AnnouncementsGroup(api.Group("/announcement"))
	announcements.CommentGroup(api.Group("/comment"))
	schedule.ScheduleGroup(api.Group("/schedule"))
	schedule.ExamsGroup(api.Group("/exams"))
	schedule.OverrideGroup(api.Group("/override"))
	api.GET("/room/:id", schedule.Room)
	api.POST("/authenticate", authentication.Authenticate)
	api.POST("/validate", authentication.AuthenticateToken)
	api.POST("/invalidate", authentication.Invalidate)
	api.POST("/register/:id", authentication.Validate(authentication.Register))
	api.POST("/unregister/:id", authentication.Validate(authentication.Unregister))
	api.GET("/rooms", schedule.Rooms)
	api.GET("/users", authentication.Validate(authentication.Users))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(types.NewContext(c, plugin, psql, redisClient, authUsers, tokens, config.Secret))
		}
	})

	//For CORS to work, please define the EXACT link that you will use!!!
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

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

	e.Server.Addr = config.Hostname + ":" + strconv.Itoa(config.Port)
	e.Logger.Fatal(graceful.ListenAndServeTLS(e.Server, "moniteur.crt", "moniteur.key", 5*time.Second))
}
