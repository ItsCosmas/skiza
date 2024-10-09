package controllers

import (
	"net/http"
	"strconv"

	"math/rand/v2"

	"github.com/gofiber/fiber/v2"
)

// ListenerController godoc
//
//	@Summary		Listen for a callback
//	@Description	Listen for a callback
//	@Tags			Listener
//	@Success		200	{object}	Response
//	@Router			/listener [post]
func ListenerController(c *fiber.Ctx, broadcast chan<- string) error {
	response := HTTPResponse(http.StatusOK, "Callback Received", "Callback Body")
	// Broadcast a message to all WebSocket clients
	broadcast <- "A callback has been received: " + strconv.Itoa(rand.IntN(100))
	return c.JSON(response)
}
