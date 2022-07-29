package route

import (
	"github.com/gofiber/fiber/v2"
	"jwt/controller"
)

func Setup(app *fiber.App) {
	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)
	app.Post("/logout", controller.Logout)

	app.Get("/user", controller.User)
}
