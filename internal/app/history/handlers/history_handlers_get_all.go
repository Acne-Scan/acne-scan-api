package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (history *HistoryHandlersImpl) GetAll(c *fiber.Ctx) error {
	id := c.Locals("user_id").(string)

	data, err := history.service.GetAll(id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "history not found", err)
		}
		return response.InternalServerError(c, "failed to get history, something happen", err.Error())
	}

	return response.StatusOk(c, http.StatusOK, "success", data)
}
