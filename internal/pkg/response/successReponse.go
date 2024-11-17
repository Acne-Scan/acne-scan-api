package response

import (
	"acne-scan-api/internal/model/web"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SuccessResponse(c *fiber.Ctx, code int, message string, data any) error {
	return c.JSON(web.SuccessResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func StatusCreated(c *fiber.Ctx, code int, message string, data any) error {
	return ErrorResponse(c, http.StatusCreated, message, data)
}

func StatusOk(c *fiber.Ctx, code int, message string, data any) error {
	return ErrorResponse(c, http.StatusOK, message, data)
}
