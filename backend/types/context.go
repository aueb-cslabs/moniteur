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
	Secret      string
}

func NewContext(context echo.Context, plugin Plugin, db *gorm.DB, redis *redis.Client, auth *redis.Client, secret string) *Context {
	return &Context{Context: context, plugin: plugin, DB: db, RedisClient: redis, AuthUsers: auth, Secret: secret}
}

func (c Context) Plugin() Plugin {
	return c.plugin
}
