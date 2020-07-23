// +build linux darwin

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

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
	cert := flag.String("cert", "moniteur.crt", "Public certificate file. Must end with .crt")
	key := flag.String("key", "moniteur.key", "Private key. Must end .key")
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

	psql := databases.InitializePostgres()
	redisClient, authUsers, tokens := databases.InitializeRedis(config.RedisConfig)

	plugin.Initialize(config.ExamsLink)
	rest.Initialize(calendarInfo)

	e := echo.New()

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
			return next(types.NewContext(c, plugin, psql, redisClient, authUsers, tokens, os.Getenv("JWT_SECRET")))
		}
	})

	//For CORS to work, please define the EXACT link that you will use!!!
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Load the certificates
	certificate, err := tls.LoadX509KeyPair(*cert, *key)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{}
	tlsConfig.Certificates = make([]tls.Certificate, 1)
	tlsConfig.Certificates[0] = certificate

	// Create a custom server with Read and Write timeouts
	s := &http.Server{
		Addr:         config.Hostname + ":" + strconv.Itoa(config.Port),
		TLSConfig:    tlsConfig,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	go func() {
		if err := e.StartServer(s); err != nil {
			e.Logger.Fatal("Shutting down the server! " + err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
