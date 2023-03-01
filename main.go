package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ionnotion/go-mongo-practice/configs"
	"github.com/ionnotion/go-mongo-practice/models"
)
func main () {
	app := fiber.New()

	configs.ConnectDatabase()

	app.Get("/",func(c *fiber.Ctx) error {
		return c.Status(200).JSON(models.Response{Message:"Hello World!"})
	})

	app.Listen(":8080")
}