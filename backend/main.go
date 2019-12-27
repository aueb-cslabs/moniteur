// +build linux darwin

package main

import "C"
import (
	"github.com/aueb-cslabs/moniteur/backend/rest"
	"github.com/aueb-cslabs/moniteur/backend/types"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/url"
	"strconv"
	"strings"
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
	calendar, err := types.LoadCalendar("config/calendar.yml")
	if err != nil {
		log.Panic(err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Addr,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.DB,
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Panic(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	plugin.Initialize(config.ExamsLink)
	rest.Initialize(config.Secret, calendar, config.AuthorizedUsers, *redisClient)

	api := e.Group("/api")

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			return next(types.NewContext(c, plugin))
		}
	})

	//For CORS to work, please define the EXACT link that you will use!!!
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080", "http://localhost:8086"},
	}))

	rest.CalendarGroup(api.Group("/calendarInfo"))
	rest.AnnouncementsGroup(api.Group("/announcement"))
	rest.CommentGroup(api.Group("/comment"))
	rest.ScheduleGroup(api.Group("/schedule"))
	rest.ExamsGroup(api.Group("/exams"))
	api.GET("/room/:id", rest.Room)
	api.POST("/authenticate", rest.Authenticate)
	api.POST("/validate", rest.AuthenticateToken)
	api.POST("/invalidate", rest.Invalidate)
	api.POST("/register/:id", rest.Validate(rest.Register))
	api.POST("/unregister/:id", rest.Validate(rest.Unregister))
	api.GET("/rooms", rest.Rooms)
	api.GET("/users", rest.Validate(rest.Users))

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
