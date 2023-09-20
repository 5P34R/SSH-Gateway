package server

import (
	"ssh-gateway/server/controllers"

	"github.com/gofiber/fiber/v2"
)

func AppRouter(app *fiber.App) {
	app.Get("/", controllers.HomeController)
	app.Get("/delete/:id", controllers.StopContainer)
}
