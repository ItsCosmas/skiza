package handler

import (
	"github.com/gofiber/fiber/v2"
)

// HomeHandler godoc
//
//	@Summary		Return a welcome message
//	@Description	Return a welcome message
//	@Tags			Home
//	@Success		200	{object}	Response
//	@Router			/home [get]
func HomeHandler(c *fiber.Ctx) error {
	response := HTTPResponse(fiber.StatusOK, "Success", "Welcome Home")
	return c.JSON(response)
}
