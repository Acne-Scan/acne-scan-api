package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (pr *ProductRecommendationHandlerImpl) Delete(c *fiber.Ctx) error{
	idparam := c.Params("id")
	if idparam=="" {
		return response.BadRequest(c, "invalid product recommendation id", nil)
	}

	err := pr.productRecommendationService.Delete(idparam)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "recommendation not found", err)
		}
		return response.InternalServerError(c, "failed to delete recommendation, something happen", err.Error())
	}

	return response.StatusOk(c,200,"success deleted recommendation",nil)
}