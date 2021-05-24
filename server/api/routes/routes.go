package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	// Controllers
	ctl "skiza/api/controllers"
)

// SetupRoutes setups router
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Use("/docs", swagger.Handler)

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Skiza Callback Listener",
		})
	})

	v1.Get("/home", ctl.HomeController)
	v1.Post("/listener", ctl.ListenerController)
}
