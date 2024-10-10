package routes

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"

	// Controllers
	ctl "skiza/api/controllers"

	// Swagger
	_ "skiza/docs" // Swagger Docs

	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/swagger"
)

type client struct {
	isClosing bool
	mu        sync.Mutex
}

var clients = make(map[*websocket.Conn]*client)
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)

func runHub() {
	for {
		select {
		case connection := <-register:
			clients[connection] = &client{}
			log.Println("Connection Registered")

		case message := <-broadcast:
			log.Println("Message Received:", message)
			// Send the message to all clients
			for connection, c := range clients {
				go func(connection *websocket.Conn, c *client) { // send to each client in parallel so we don't block on a slow client
					c.mu.Lock()
					defer c.mu.Unlock()
					if c.isClosing {
						return
					}
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.isClosing = true
						log.Println("write error:", err)

						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-unregister:
			// Remove the client from the hub
			delete(clients, connection)

			log.Println("Connection Unregistered")
		}
	}
}

// setup router
func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	v1 := api.Group("/v1")

	// Swagger
	v1.Get("/docs/*", swagger.HandlerDefault)

	// v1.Use(func(c *fiber.Ctx) error {
	// 	if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
	// 		return c.Next()
	// 	}
	// 	return c.SendStatus(fiber.StatusUpgradeRequired)
	// })

	go runHub()

	// Upgraded websocket request
	v1.Get("/ws", websocket.New(func(c *websocket.Conn) {
		fmt.Println(c.Locals("Host")) // "Localhost:3000"
		defer func() {
			unregister <- c
			c.Close()
		}()

		// Register the client
		register <- c

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				// Broadcast the received message
				broadcast <- string(message)
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))

	v1.Get("/home", ctl.HomeController)
	v1.Post("/listener", func(c *fiber.Ctx) error {
		return ctl.ListenerController(c, broadcast)
	})
}
