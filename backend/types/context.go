package types

import (
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	plugin      Plugin
	Calendar    *Calendar
	DB          *gorm.DB
	RedisClient *redis.Client
	AuthUsers   *redis.Client
}

func NewContext(context echo.Context, plugin Plugin, db *gorm.DB, redis *redis.Client, auth *redis.Client) *Context {
	return &Context{Context: context, plugin: plugin, DB: db, RedisClient: redis, AuthUsers: auth}
}

func (c Context) Plugin() Plugin {
	return c.plugin
}
