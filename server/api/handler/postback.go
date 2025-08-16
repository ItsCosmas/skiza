package handler

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

// ExternalResponse is the structure of the json message
type ExternalResponse struct {
	ResponseBody string `json:"responseBody"`
	SourceIP     string `json:"sourceIP"`
}

// PostbackListener godoc
//
//	@Summary		Listen for a callback
//	@Description	Listen for a callback
//	@Tags			Listener
//	@Success		200	{object}	Response
//	@Router			/listener [post]
func PostbackListener(c *fiber.Ctx, broadcaster func(string)) error {
	httpResponse := HTTPResponse(fiber.StatusOK, "Callback Received", "Callback Body")

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
		return c.Status(fiber.StatusInternalServerError).SendString("Error encoding response")
	}

	// Broadcast the ExternalResponse in a goroutine
	go func() {
		// Broadcast the ExternalResponse to all WebSocket clients
		log.Println("Broadcast Sent")
		broadcaster(string(jsonResponse))
	}()

	// Acknowledge the HTTP request
	log.Println("HTTP Response Sent")
	return c.JSON(httpResponse)
}
