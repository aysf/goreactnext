package routes

import (
	"github.com/aysf/goreactnext/src/controllers"
	"github.com/aysf/goreactnext/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	admin := api.Group("admin")

	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)

	auth := admin.Use(middlewares.IsAuthenticated)
	auth.Get("/user", controllers.User)
	auth.Get("/logout", controllers.Logout)
	auth.Put("/user/info", controllers.UpdateInfo)
	auth.Put("/user/password", controllers.UpdatePassword)
	auth.Get("/ambassadors", controllers.Ambassadors)
}
