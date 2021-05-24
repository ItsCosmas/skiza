package app

import (
	"fmt"

	// Configs
	cfg "skiza/api/configs"

	// Swagger
	docs "skiza/api/docs" // Swagger Docs

	// routes
	"skiza/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// Run starts the app
// @title Skiza API
// @version 1.0
// @description Skiza is a Callback Listener and Pusher.
// @termsOfService http://swagger.io/terms/
// @contact.name Cozy
// @contact.url https://github.com/ItsCosmas
// @contact.email devcosmas@gmail.com
// @license.name MIT
// @license.url https://github.com/ItsCosmas/skiza/blob/master/LICENSE
// @BasePath /api/v1
func Run() {
	app := fiber.New()

	/*
		====== Setup Configs ============
	*/

	cfg.LoadConfig()
	config := cfg.GetConfig()

	/*
		============ Set Up Middlewares ============
	*/

	// Default Log Middleware
	app.Use(logger.New())

	// Recovery Middleware
	app.Use(recover.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	/*
		============ Set Up Routes ============
	*/
	routes.SetupRoutes(app)

	/*
		============ Setup Swagger ===============
	*/

	// FIXME, In Production, Port Should not be added to the Swagger Host
	docs.SwaggerInfo.Host = config.Host + ":" + config.Port

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
