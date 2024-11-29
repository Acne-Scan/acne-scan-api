package service

import (
	"acne-scan-api/internal/model/web"
	conversion "acne-scan-api/internal/pkg/conversion/request"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (articleService *ArticleServiceImpl) Create(request web.ArticleCreateRequest, image *multipart.FileHeader, c *fiber.Ctx) error {
	var err error

	err = articleService.Validator.Struct(request)
	if err != nil {
		return err
	}

	article := conversion.ArticleCreateRequestToArticleModel(request)

	imageUrl, err := articleService.BucketUploder.Uploader(c, image)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//assign image value
	article.Image = imageUrl

	wib, err := time.LoadLocation("Asia/Jakarta") // WIB (UTC+7)
	if err != nil {
		return fmt.Errorf("error loading WIB location: %s", err.Error())
	}

	createdAt := time.Now().In(wib)
	article.CreatedAt = createdAt
	article.UpdatedAt = createdAt

	err = articleService.ArticleRepository.Create(article)
	if err != nil {
		return fmt.Errorf("error when creating article %s", err.Error())
	}

	return nil
}
