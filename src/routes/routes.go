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

	adminAuth := admin.Use(middlewares.IsAuthenticated)
	adminAuth.Get("user", controllers.User)
	adminAuth.Get("logout", controllers.Logout)
	adminAuth.Put("user/info", controllers.UpdateInfo)
	adminAuth.Put("user/password", controllers.UpdatePassword)
	adminAuth.Get("ambassadors", controllers.Ambassadors)
	adminAuth.Get("products", controllers.Products)
	adminAuth.Post("products", controllers.CreateProduct)
	adminAuth.Get("products/:id", controllers.GetProduct)
	adminAuth.Put("products/:id", controllers.UpdateProduct)
	adminAuth.Delete("products/:id", controllers.DeleteProduct)
	adminAuth.Get("users/:id/links", controllers.Link)
	adminAuth.Get("orders", controllers.Orders)

	ambassador := api.Group("ambassador")
	ambassador.Post("login", controllers.Login)
	ambassador.Post("register", controllers.Register)
	ambassador.Get("products/frontend", controllers.ProductFrontend)
	ambassador.Get("products/backend", controllers.ProductBackend)

	ambassadorAuth := ambassador.Use(middlewares.IsAuthenticated)
	ambassadorAuth.Get("user", controllers.User)
	ambassadorAuth.Get("logout", controllers.Logout)
	ambassadorAuth.Put("user/info", controllers.UpdateInfo)
	ambassadorAuth.Put("user/password", controllers.UpdatePassword)
	ambassadorAuth.Post("links", controllers.CreateLinks)
	ambassadorAuth.Get("stats", controllers.Stats)
	ambassadorAuth.Get("ranking", controllers.Ranking)

}
