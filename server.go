package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	database "ohas-api.com/v1/databases"
	"ohas-api.com/v1/routes"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	database.Connect()

	routes.Routing(e)
	e.Logger.Fatal(e.Start(":8000"))
}
