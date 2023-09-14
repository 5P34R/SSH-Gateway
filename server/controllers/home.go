package controllers

import (
	"ssh-gateway/gateway"

	"github.com/gofiber/fiber/v2"
)

type Container struct {
	Name        string
	ActiveUsers int
	UpTime      string
	Health      string
}

func HomeController(c *fiber.Ctx) error {

	containerCount := gateway.GetNumberOfContainer()

	return c.Render("index", fiber.Map{
		"Title":          "Hello World",
		"containerCount": containerCount,
	})
}
