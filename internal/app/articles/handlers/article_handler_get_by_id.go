package handlers

import (
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (articleHandler *ArticleHandlerImpl) GetById(c *fiber.Ctx) error {
	idparam := c.Params("id")
	if idparam=="" {
		return response.BadRequest(c, "invalid article id", nil)
	}

	data, err := articleHandler.ArticleService.GetById(idparam)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "article not found", err)
		}
		return response.InternalServerError(c, "failed to get article, something happen", err.Error())
	}

	return response.StatusOk(c, http.StatusOK, "success", data)

}
