package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Response object as HTTP response
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"body"`
}

// ErrorBody object
type ErrorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// FiberErrorResponse object
type FiberErrorResponse struct {
	Error []*ErrorBody `json:"errors"`
}

// ErrorResponse object
type ErrorResponse struct {
	Error []*Response `json:"errors"`
}

// HTTPResponse normalize HTTP Response format
func HTTPResponse(httpCode int, message string, data interface{}) *Response {
	return &Response{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
}

// HTTPFiberErrorResponse normalizes error responses
func HTTPFiberErrorResponse(errorObj []*fiber.Error) *FiberErrorResponse {
	// Convert fiber.Error to ErrorBody
	// This fixes issues with swagger auto generated docs not identify fiber.Error type
	errorSlice := make([]*ErrorBody, 0, len(errorObj))
	for _, err := range errorObj {
		errorSlice = append(errorSlice, mapToErrorOutput(err))
	}

	return &FiberErrorResponse{
		Error: errorSlice,
	}
}

// HTTPErrorResponse normalizes error responses
func HTTPErrorResponse(errorObj []*Response) *ErrorResponse {
	return &ErrorResponse{
		Error: errorObj,
	}
}

// CreateErrorList creates a new error list (replaces global errorList)
func CreateErrorList() []*Response {
	return make([]*Response, 0)
}

// AddError adds an error to the error list and returns the updated list
func AddError(errorList []*Response, code int, message string, data interface{}) []*Response {
	return append(errorList, &Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// SendHTTPError is a generic helper to send single error responses
func SendHTTPError(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	errorList := CreateErrorList()
	errorList = AddError(errorList, statusCode, message, data)
	return c.Status(statusCode).JSON(HTTPErrorResponse(errorList))
}

// SendHTTPErrors is a generic helper to send multiple error responses
func SendHTTPErrors(c *fiber.Ctx, statusCode int, errorList []*Response) error {
	return c.Status(statusCode).JSON(HTTPErrorResponse(errorList))
}

// Example usage in your handlers:
//   return SendHTTPError(c, http.StatusBadRequest, "Error message", nil)
//
// Or for multiple errors:
//   errorList := CreateErrorList()
//   errorList = AddError(errorList, http.StatusBadRequest, "First error", nil)
//   errorList = AddError(errorList, http.StatusBadRequest, "Second error", nil)
//   return SendHTTPErrors(c, http.StatusBadRequest, errorList)

// ==================================== //
// Private Method
func mapToErrorOutput(e *fiber.Error) *ErrorBody {
	return &ErrorBody{
		Code:    e.Code,
		Message: e.Message,
	}
}
