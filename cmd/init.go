package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initHttpServer() *echo.Echo {
	var srv = echo.New()
	srv.HideBanner = true
	srv.Use(middleware.Logger())
	srv.Use(middleware.Recover())
	initHttpHandlers(srv)
	return srv
}
