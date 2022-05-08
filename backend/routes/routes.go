package routes

import (
	"backend/auths"
	"backend/hello"
	"backend/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

func Get() *echo.Echo {

	log.Println("Setting up routes")
	e := echo.New()
	e.HideBanner = true
	e.Static("/", "ui")
	echo.NotFoundHandler = func(c echo.Context) error {
		return c.File("ui/index.html")
	}
	e.Use(middleware.CORSWithConfig(middlewares.CORSConfig()))
	e.GET("/auth/login", auths.LoginHandler)
	e.GET("/auth/callback", auths.AuthCodeHandler)
	// ==========================v1 APIs=======================
	v1 := e.Group("/api/v1", middleware.JWTWithConfig(middlewares.JWTConfig()))
	v1.GET("/employees", hello.GetHandler)

	return e

}
