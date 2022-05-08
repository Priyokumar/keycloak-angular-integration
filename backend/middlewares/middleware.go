package middlewares

import (
	"backend/auths"
	"backend/configs"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func CORSConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Accept", "Cache-Control", "Content-Type", "X-Requested-With", "Authorization"},
		AllowMethods: []string{"POST", "PUT", "GET", "OPTIONS", "PATCH"},
	}
}

func JWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		TokenLookup: "header:Authorization,cookie:" + configs.Configs.Cookie,
		ParseTokenFunc: func(token string, ctx echo.Context) (interface{}, error) {
			_, claims, err := auths.VerifyToken(token)
			if err != nil {
				logrus.Println("Error while verifying jwt token in middleware")
				logrus.Println(err.Error())
				http.Redirect(ctx.Response().Writer, ctx.Request(), "/", http.StatusPermanentRedirect)
			}
			user := auths.ClaimsToUser(claims)
			return user, nil
		},
		AuthScheme: "Bearer",
		ErrorHandlerWithContext: func(err error, ctx echo.Context) error {
			logrus.Println("Error while checking jwt token in middleware")
			logrus.Println(err.Error())
			http.Redirect(ctx.Response().Writer, ctx.Request(), "/", http.StatusPermanentRedirect)
			return nil
		},
	}
}
