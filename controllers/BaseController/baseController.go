package basecontroller

import (
	"github.com/labstack/echo/v4"
	"ohas-api.com/v1/helpers"
)

type Hellaw struct {
	Message string
}

func WelcomeMessage(c echo.Context) error {

	res := Hellaw{
		Message: "Hellow World",
	}
	return helpers.ResponseSuccess(c, res)
}
