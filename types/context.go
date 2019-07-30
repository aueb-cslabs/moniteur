package types

import "github.com/labstack/echo"

type Context struct {
	echo.Context
	plugin   Plugin
	Calendar *Calendar
}

func NewContext(context echo.Context, plugin Plugin, calendar *Calendar) *Context {
	return &Context{Context: context, plugin: plugin, Calendar: calendar}
}

func (c Context) Plugin() Plugin {
	return c.plugin
}
