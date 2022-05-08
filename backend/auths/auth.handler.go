package auths

import (
	"backend/configs"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func LoginHandler(ctx echo.Context) error {

	log.Println("Inside Login Handler")

	redirectPath := getOauth2Config().AuthCodeURL(configs.Configs.AuthState)

	http.Redirect(ctx.Response().Writer, ctx.Request(), redirectPath, http.StatusTemporaryRedirect)

	return nil
}

func AuthCodeHandler(ctx echo.Context) error {

	if ctx.Request().URL.Query().Get("state") != configs.Configs.AuthState {
		http.Error(ctx.Response().Writer, "state did not match", http.StatusBadRequest)
		return nil
	}
	oauth2Token, err := getOauth2Config().Exchange(context.Background(), ctx.Request().URL.Query().Get("code"))
	if err != nil {
		http.Error(ctx.Response().Writer, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return nil
	}
	accessToken, _, err := VerifyToken(oauth2Token.AccessToken)
	if err != nil {
		http.Error(ctx.Response().Writer, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return nil
	}
	accessTokenCookie := http.Cookie{
		Name:     configs.Configs.Cookie,
		Value:    oauth2Token.AccessToken,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		Expires:  accessToken.Expiry,
	}
	http.SetCookie(ctx.Response().Writer, &accessTokenCookie)
	http.Redirect(ctx.Response().Writer, ctx.Request(), "/", http.StatusTemporaryRedirect)
	return nil
}
