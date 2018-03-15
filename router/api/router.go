package api

import (
	"github.com/labstack/echo"
	. "echo-web/conf"
	mw "github.com/labstack/echo/middleware"
	"mm-echo-web/module/session"
	"mm-echo-web/module/cache"
)

func Routers() *echo.Echo {
	e := echo.New()
	e.Use(NewContext())

	if Conf.ReleaseMode {
		e.Debug = false
	}

	e.Logger.SetPrefix("api")
	e.Logger.SetLevel(GetLogLvl())

	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		TokenLookup : "form:X-XSRF-TOKEN",
	}))

	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level : 5,
	}))

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Static("/favicon.ico", "./assets/img/favicon.ico")	

	e.Use(session.Session())

	e.Use(cache.Cache())

	e.GET("/login", UserLoginHandler)

	r := e.Group("")

	r.Use(mw.JWTWithConfig(mw.JWTConfig{
		SigningKey : []byte("secret"),
		ContextKey : "_user",
		TokenLookup : "header:" + echo.HeaderAuthorization,
	}))

	r.GET("/user", UserHandler)
	return e
}