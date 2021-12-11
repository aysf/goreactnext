package main

import (
	"github.com/aysf/goreactnext/src/database"
	"github.com/aysf/goreactnext/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()
	database.Automigrate()

	app := fiber.New()

	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(":) Hello, World! fd :)")
	})

	app.Listen(":8080")
}
