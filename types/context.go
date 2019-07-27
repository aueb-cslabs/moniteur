package types

import "github.com/labstack/echo"

type Context struct {
	echo.Context
	plugin Plugin
	Rooms  *RoomMap
}

func NewContext(context echo.Context, plugin Plugin, rooms *RoomMap) *Context {
	return &Context{Context: context, plugin: plugin, Rooms: rooms}
}

func (c Context) Plugin() Plugin {
	return c.plugin
}
