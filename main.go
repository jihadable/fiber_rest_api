package main

import (
	"fiber_rest_api/database"
	"fiber_rest_api/route"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	env := viper.New()
	env.SetConfigFile(".env")
	env.AddConfigPath("../")

	err := env.ReadInConfig()
	if err != nil {
		panic(err)
	}

	port := env.GetString("PORT")

	app := fiber.New()
	db := database.GetConnection()

	api := app.Group("/api")

	route.ProductRoute(api, db)

	app.Listen("localhost:" + port)
}
