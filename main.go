package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	database "ohas-api.com/v1/databases"
	"ohas-api.com/v1/routes"
)

func main() {
	e := echo.New()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error on loading .env file")
	}

	database.Connect()

	routes.Routing(e)
	e.Logger.Fatal(e.Start(":8000"))
}
