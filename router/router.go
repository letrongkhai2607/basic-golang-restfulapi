package router

import (
	"example.com/m/v2/controller"
	"example.com/m/v2/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)

	// Transfer
	transaction := api.Group("/transaction")
	transaction.Get("/transfer" , controller.GetAllTransferHistory)
	transaction.Get("/transfer/:id" , controller.GetTransferHistoryByUserId)
	transaction.Post("/transfer" , controller.Transfer)

	transaction.Get("/recharge" , controller.GetAllRechargeHistory)
	transaction.Get("/recharge/:id" , controller.GetRechargeHistoryByUserId)
	transaction.Post("/recharge" , controller.Recharge)

	// User
	user := api.Group("/users")
	user.Get("/", controller.GetAllUsers)
	user.Get("/:id", controller.GetUserById)
	user.Post("/", controller.CreateNewUser)
	user.Patch("/:id", controller.UpdateUser)
	user.Delete("/:id", controller.DeleteUser)
	// user.Patch("/:id", middleware.Protected(), controller.UpdateUser)
	// user.Delete("/:id", middleware.Protected(), controller.DeleteUser)

	// Product
	product := api.Group("/products")
	product.Get("/", controller.GetAllProducts)
	product.Get("/:id", controller.GetProduct)
	product.Post("/", controller.CreateProduct)
	product.Delete("/:id", middleware.Protected(), controller.DeleteProduct)
}