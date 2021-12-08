package main

import (
	"github.com/aysf/goreactnext/src/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	database.Automigrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(":) Hello, World! fd :) :)")
	})

	app.Listen(":8080")
}
