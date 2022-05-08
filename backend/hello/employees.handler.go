package hello

import "github.com/labstack/echo/v4"

func GetHandler(ctx echo.Context) error {

	ctx.JSON(200, []string{"Apang", "Budu", "Crack"})

	return nil

}
