// +build linux darwin

package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tylerb/graceful"

	"github.com/aueb-cslabs/moniteur/backend/databases"
	"github.com/aueb-cslabs/moniteur/backend/rest"
	"github.com/aueb-cslabs/moniteur/backend/rest/announcements"
	"github.com/aueb-cslabs/moniteur/backend/rest/authentication"
	"github.com/aueb-cslabs/moniteur/backend/rest/calendar"
	"github.com/aueb-cslabs/moniteur/backend/rest/schedule"
	"github.com/aueb-cslabs/moniteur/backend/types"
)

func main() {
	configFile := flag.String("config", "config.yml", "Specifies the configuration file to be read")
	pluginFIle := flag.String("plugin", "", "Specifies the plugin file to bread (Required)")
	calendarFile := flag.String("calendar", "calendar.yml", "Specifies the calendar file to be read")
	help := flag.Bool("help", false, "Displays this message")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *pluginFIle == "" {
		fmt.Println("Error! Please specify a plugin file!")
		os.Exit(-1)
	}

	config, err := types.LoadConfiguration(*configFile)
	if err != nil {
		log.Panic(err)
	}
	plugin, err := types.LoadPlugin(*pluginFIle)
	if err != nil {
		log.Panic(err)
	}
	calendarInfo, err := types.LoadCalendar(*calendarFile)
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
