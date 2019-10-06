package types

import "github.com/labstack/echo"

type Context struct {
	echo.Context
	plugin          Plugin
	Calendar        *Calendar
	AuthorizedUsers map[string]string
}

func NewContext(context echo.Context, plugin Plugin, calendar *Calendar, authorizedUsers map[string]string) *Context {
	return &Context{Context: context, plugin: plugin, Calendar: calendar, AuthorizedUsers: authorizedUsers}
}

func (c Context) Plugin() Plugin {
	return c.plugin
}
