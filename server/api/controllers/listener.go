package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"skiza/api/schema"

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
	httpResponse := HTTPResponse(http.StatusOK, "Callback Received", "Callback Body")

	// Create an ExternalResponse object
	responseBody := string(c.Body())
	sourceIP := c.IP() // Get the source IP

	response := schema.ExternalResponse{
		ResponseBody: responseBody,
		SourceIP:     sourceIP,
	}

	// Convert ExternalResponse to JSON (byte slice)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error encoding response")
	}

	// return c.JSON(httpResponse)

	c.JSON(httpResponse)

	// Broadcast the ExternalResponse in a goroutine
	go func() {
		// Optionally log the body of the request
		fmt.Println(responseBody)

		// Broadcast the ExternalResponse to all WebSocket clients
		broadcast <- string(jsonResponse)
	}()

	return nil // Return nil since we already sent the response
}
