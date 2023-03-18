package routes

import (
	"github.com/labstack/echo/v4"
	authcontroller "ohas-api.com/v2/controllers/AuthController"
	basecontroller "ohas-api.com/v2/controllers/BaseController"
)

func Routing(e *echo.Echo) {
	apiRoutes := e.Group("/api")

	apiRoutes.GET("/Users", basecontroller.WelcomeMessage)
	apiRoutes.POST("/register", authcontroller.Register)
	apiRoutes.POST("/login", authcontroller.Login)
}
