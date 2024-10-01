package routes

import (
	"github.com/gofiber/fiber/v2"

	// Controllers
	ctl "skiza/api/controllers"

	// Swagger
	_ "skiza/docs" // Swagger Docs

	"github.com/gofiber/swagger"
)

// setup router
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// Swagger
	v1.Get("/docs/*", swagger.HandlerDefault)

	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Skiza Callback Listener",
		})
	})

	v1.Get("/home", ctl.HomeController)
	v1.Post("/listener", ctl.ListenerController)
}
