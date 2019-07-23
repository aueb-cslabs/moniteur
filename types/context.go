package types

import "github.com/labstack/echo"

type Context struct {
	echo.Context
	plugin Plugin
}

func NewContext(context echo.Context, plugin Plugin) *Context {
	return &Context{Context: context, plugin: plugin}
}

func (c Context) Plugin() Plugin {
	return c.plugin
}