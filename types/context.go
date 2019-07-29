package types

import "github.com/labstack/echo"

type Context struct {
	echo.Context
	plugin   Plugin
	Rooms    *RoomMap
	Calendar *Calendar
}

func NewContext(context echo.Context, plugin Plugin, rooms *RoomMap, calendar *Calendar) *Context {
	return &Context{Context: context, plugin: plugin, Rooms: rooms, Calendar: calendar}
}

func (c Context) Plugin() Plugin {
	return c.plugin
}
