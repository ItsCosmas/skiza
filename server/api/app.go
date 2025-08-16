package app

import (
	"fmt"
	"log"

	// Configs
	cfg "skiza/api/config"

	// Handlers
	"skiza/api/handler"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Run() {
	app := fiber.New()

	// Configs
	cfg.LoadConfig()
	config := cfg.GetConfig()

	// Middlewares
	// Default Log Middleware
	app.Use(logger.New())
	// Recovery Middleware
	app.Use(recover.New())
	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Start WebSocket hub once
	handler.InitHub()

	// Routes
	api := app.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/home", handler.HomeHandler)
	v1.Post("/listener", func(c *fiber.Ctx) error {
		return handler.PostbackListener(c, handler.BroadcastMessage)
	})

	// WebSocket upgrade middleware
	v1.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// WebSocket route
	v1.Get("/ws", websocket.New(handler.Stream))

	// Run the app and listen on given port
	port := fmt.Sprintf(":%s", config.Port)
	log.Fatal(app.Listen(port))
}
