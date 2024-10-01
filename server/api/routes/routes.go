package routes

import (
	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"

	// Controllers
	ctl "skiza/api/controllers"
)

// SetupRoutes setups router
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	v1.Use("/docs", swagger.HandlerDefault)

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Skiza Callback Listener",
		})
	})

	v1.Get("/home", ctl.HomeController)
	v1.Post("/listener", ctl.ListenerController)
}
