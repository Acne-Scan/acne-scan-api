package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (pr *ProductRecommendationHandlerImpl) GetById(c *fiber.Ctx)error{
	idparam := c.Params("id")
	if idparam=="" {
		return response.BadRequest(c, "invalid product recommendation id", nil)
	}

	data, err := pr.productRecommendationService.GetById(idparam)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "product recommendation not found", err)
		}
		return response.InternalServerError(c, "failed to get product recommendation, something happen", err.Error())
	}

	return response.StatusOk(c, http.StatusOK, "success", data)
}