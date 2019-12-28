package types

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Context struct {
	echo.Context
	plugin   Plugin
	Calendar *Calendar
	DB       *gorm.DB
}

func NewContext(context echo.Context, plugin Plugin, db *gorm.DB) *Context {
	return &Context{Context: context, plugin: plugin, DB: db}
}

func (c Context) Plugin() Plugin {
	return c.plugin
}
