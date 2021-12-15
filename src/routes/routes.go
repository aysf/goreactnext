package routes

import (
	"github.com/aysf/goreactnext/src/controllers"
	"github.com/aysf/goreactnext/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api")

	admin := api.Group("admin")

	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	auth := admin.Use(middlewares.IsAuthenticated)
	auth.Get("user", controllers.User)
	auth.Get("logout", controllers.Logout)
	auth.Put("user/info", controllers.UpdateInfo)
	auth.Put("user/password", controllers.UpdatePassword)
	auth.Get("ambassadors", controllers.Ambassadors)

	// products
	auth.Get("products", controllers.Products)
	auth.Post("products", controllers.CreateProduct)
	auth.Get("products/:id", controllers.GetProduct)
	auth.Put("products/:id", controllers.UpdateProduct)
	auth.Delete("products/:id", controllers.DeleteProduct)

	auth.Get("users/:id/links", controllers.Link)

	// orders
	auth.Get("orders", controllers.Orders)

}
