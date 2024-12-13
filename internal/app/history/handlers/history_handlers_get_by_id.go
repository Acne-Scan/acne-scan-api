package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (history *HistoryHandlersImpl) GetById(c *fiber.Ctx) error {
	idparam := c.Params("id")
	if idparam == "" {
		return response.BadRequest(c, "invalid history users id", nil)
	}

	data, err := history.service.GetById(idparam)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "history not found", err)
		}
		if strings.Contains(err.Error(), "no result") {
			return response.StatusOk(c,http.StatusOK, "history not found",nil)
		}
		return response.InternalServerError(c, "failed to get history, something happen", err.Error())
	}

	return response.StatusOk(c, http.StatusOK, "success", data)
}
