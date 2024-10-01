package app

import (
	"fmt"

	// Configs
	cfg "skiza/api/configs"

	// routes
	"skiza/api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	app := fiber.New()

	// ====== Configs ============
	cfg.LoadConfig()
	config := cfg.GetConfig()

	// ============ Middlewares ============
	// Default Log Middleware
	app.Use(logger.New())
	// Recovery Middleware
	app.Use(recover.New())
	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// ============ Routes ============
	routes.SetupRoutes(app)

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	app.Listen(port)
}
