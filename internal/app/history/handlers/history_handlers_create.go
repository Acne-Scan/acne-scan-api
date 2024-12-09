package handlers

import (
	"acne-scan-api/internal/model/web"
	"acne-scan-api/internal/pkg/response"
	"acne-scan-api/internal/pkg/validation"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (history *HistoryHandlersImpl) Create(c *fiber.Ctx) error {
	req := new(web.HistoryRequest)

	// var Profileimage *multipart.FileHeader

	Profileimage, err := c.FormFile("user_picture")
	if err != nil {
		if Profileimage != nil {
			return err
		}
	}

	userId := c.Locals("user_id").(int)

	if err := c.BodyParser(req); err != nil {
		fmt.Println(err.Error())
		return response.BadRequest(c, "failed to bind history request", err)
	}

	ProductLinkJson, err := json.Marshal(req.ProductLink)
	if err != nil {
		log.Fatalf("Error serializing map to JSON: %v", err)
		return response.InternalServerError(c, "failed to serializing map to JSON", err.Error())
	}

	historyJson, err := json.Marshal(req.Image)
	if err != nil {
		log.Fatalf("Error serializing map to JSON: %v", err)
		return response.InternalServerError(c, "failed to serializing map to JSON", err.Error())
	}

	err = history.service.Create(c, req, historyJson, Profileimage, ProductLinkJson, userId)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(c, err)
		}
		if strings.Contains(err.Error(), "FOREIGN KEY (`user_id`)") {
			return response.StatusOk(c, http.StatusOK, "Data user tidak ditemukan", nil)
		}
		return response.InternalServerError(c, "failed to create history, something happen", err.Error())
	}

	return response.StatusCreated(c, http.StatusCreated, "success to create history", nil)
}
