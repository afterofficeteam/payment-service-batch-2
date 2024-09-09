package route

import (
	paymentHandler "codebase-app/internal/module/payment/handler/rest"

	"codebase-app/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// add /api prefix to all routes
	api := app.Group("/payments")
	paymentHandler.NewpaymentHandler().Register(api)

	// health check route
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(response.Success(nil, "Server is running."))
	})

	// fallback route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(response.Error("Route not found."))
	})
}
