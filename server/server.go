package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Start(port string) {
	engine := html.New("./template", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	AppRouter(app)
	log.Fatal(app.Listen(port))
}
