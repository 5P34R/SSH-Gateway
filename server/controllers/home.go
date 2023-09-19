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
	containers := gateway.ContainerList()

	return c.Render("index", fiber.Map{
		"Title":          "Hello World",
		"containerCount": containerCount,
		"containers":     containers,
	})
}

func DeleteContainer(c *fiber.Ctx) error {
	ID := c.Params("id")
	return c.JSON(fiber.Map{
		"container id": ID,
	})
}
