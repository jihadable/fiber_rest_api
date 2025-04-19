package main

import (
	"fiber_rest_api/database"
	"fiber_rest_api/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	db := database.GetConnection()

	api := app.Group("/api")

	route.ProductRoute(api, db)

	app.Listen("localhost:8000")
}
