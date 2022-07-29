package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"jwt/database"
	"jwt/route"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	route.Setup(app)

	listenErr := app.Listen(":3000")
	if listenErr != nil {
		panic("failed to connect to port 3000")
	}
}
