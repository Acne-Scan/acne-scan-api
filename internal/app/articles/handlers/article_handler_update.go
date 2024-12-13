package handlers

import (
	"acne-scan-api/internal/model/web"
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (articleHandler *ArticleHandlerImpl) Update(c *fiber.Ctx) error {
	idparam := c.Params("id")
	if idparam=="" {
		return response.BadRequest(c, "invalid article id", nil)
	}

	req:=new(web.ArticleUpdateRequest)
	if err := c.BodyParser(req); err != nil {
		fmt.Println(err.Error())
		return response.BadRequest(c, "failed to bind article request", err)
	}

	ifExist,err:=articleHandler.ArticleService.GetById(idparam)
	if ifExist==nil {
		return response.BadRequest(c, "cannot found article to update", err)
	}

	err=articleHandler.ArticleService.Update(req.Name,req.Description,req.Image,idparam)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return response.BadRequest(c, "cannot found article to update", err)
		}
		return response.InternalServerError(c, "failed to update article, something happen", err.Error())
	}

	return response.StatusOk(c,200,"article updated",nil)
}