package main

import (
	"github.com/aysf/goreactnext/src/database"
	"github.com/aysf/goreactnext/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()
	database.Automigrate()
	database.SetupRedis()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(":) Hello, World! fd :)")
	})

	app.Listen(":8080")
}
