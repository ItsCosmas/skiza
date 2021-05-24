package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ListenerController godoc
// @Summary Listen for a callback
// @Description Listen for a callback
// @Tags Listener
// @Success 200 {object} Response
// @Router /listener [post]
func ListenerController(c *fiber.Ctx) error {
	response := HTTPResponse(http.StatusOK, "Callback Received", "Callback Body")
	return c.JSON(response)
}
