package service

import (
	"acne-scan-api/internal/model/web"
	conversion "acne-scan-api/internal/pkg/conversion/request"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (articleService *ArticleServiceImpl) Create(request web.ArticleCreateRequest, image *multipart.FileHeader, c *fiber.Ctx) error {
	var err error

	err = articleService.Validator.Struct(request)
	if err != nil {
		return err
	}

	article := conversion.ArticleCreateRequestToArticleModel(request)

	imageUrlChan := make(chan string)
	errChan := make(chan string)

	go func() {
		imageUrl, err := articleService.BucketUploder.Uploader(c, image)
		if err != nil {
			errChan <- err.Error()
			return
		}

		imageUrlChan <- imageUrl
	}()

	select {
	case imageUrl := <-imageUrlChan:
		article.Image = imageUrl
	case err := <-errChan:
		return fmt.Errorf("error uploading image: %s", err)
	}

	wib, err := time.LoadLocation("Asia/Jakarta") // WIB (UTC+7)
	if err != nil {
		return fmt.Errorf("error loading WIB location: %s", err.Error())
	}

	//generate random id
	randomID := uuid.New().String()

	createdAt := time.Now().In(wib)
	article.CreatedAt = createdAt
	article.UpdatedAt = createdAt
	article.ArticleId = randomID

	err = articleService.ArticleRepository.Create(article)
	if err != nil {
		return fmt.Errorf("error when creating article %s", err.Error())
	}

	return nil
}
