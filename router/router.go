package router

import (
	"context"
	"net/url"
	"os"
	"os/signal"
	"time"
	
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	. "mm-echo-web/conf"
	mw "github.com/labstack/echo/middleware"
	"mm-echo-web/router/api"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)


func InitRoutes() map[string]*Host {
	// Hosts
	hosts := make(map[string]*Host)

	//hosts[Conf.Server.DomainWeb] = &Host{web.Routers()}
	hosts[Conf.Server.DomainApi] = &Host{api.Routers()}
	//hosts[Conf.Server.DomainSocket] = &Host{socket.Routers()}

	return hosts
}

func RunSubdomains(confFilePath string){
	if err := InitConfig(confFilePath); err != nil {
		log.Panic(err)
	}
	
	log.SetLevel(GetLogLvl())

	e := echo.New()
	e.Pre(mw.RemoveTrailingSlash())

	e.Logger.SetLevel(GetLogLvl())

	e.Use(mw.SecureWithConfig(mw.DefaultSecureConfig))
	mw.MethodOverride()

	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"http://" + Conf.Server.DomainWeb, "http://" + Conf.Server.DomainApi},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
	}))
	
	hosts := InitRoutes()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		u, _err := url.Parse(c.Scheme() + "://" + req.Host)
		if _err != nil {
			e.Logger.Errorf("Request URL parse error:%v", _err)
		}

		host := hosts[u.Hostname()]
		if host == nil {
			e.Logger.Info("Host not found")
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

	if !Conf.Server.Graceful {
		e.Logger.Fatal(e.Start(Conf.Server.Addr))
	} else {
		// Graceful Shutdown
		// Start server
		go func() {
			if err := e.Start(Conf.Server.Addr); err != nil {
				e.Logger.Errorf("Shutting down the server with error:%v", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}
}