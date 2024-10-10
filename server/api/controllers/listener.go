package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ExternalResponse is the structure of the json message
type ExternalResponse struct {
	ResponseBody string `json:"responseBody"`
	SourceIP     string `json:"sourceIP"`
}

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

	response := ExternalResponse{
		ResponseBody: responseBody,
		SourceIP:     sourceIP,
	}

	// Convert ExternalResponse to JSON (byte slice)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error encoding response")
	}

	// Broadcast the ExternalResponse in a goroutine
	go func() {
		// Broadcast the ExternalResponse to all WebSocket clients
		log.Println("Broadcast Sent")
		broadcast <- string(jsonResponse)
	}()

	// Immediately return the HTTP response and execute the go routine after
	log.Println("HTTP Response Sent")
	return c.JSON(httpResponse)
}
