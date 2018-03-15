package api

import (
	"github.com/labstack/echo"
	"github.com/hb-go/json"
)

func NewContext() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &Context{c}
			return h(ctx)
		}
	}
}

type Context struct {
	echo.Context
}

func (c *Context) AutoFMT(code int, i interface{}) (err error) {
	callback := c.QueryParam("jsonp")
	if len(callback) > 0 {
		c.Logger().Infof("JSONP callback func:%v", callback)
		return c.JSONP(code, callback, i)
	} else {
		return c.JSON(code, i)
	}
}

func (c *Context) CustomJSON(code int, i interface{}, f string) (err error) {
	if c.Context.Echo().Debug {
		return c.JSONPretty(code, i, " ")
	}
	b, err := json.MarshalFilter(i, f)
	if err != nil {
		return
	}
	return c.JSONBlob(code, b)
}