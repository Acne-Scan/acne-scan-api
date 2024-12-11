package service

import (
	"acne-scan-api/internal/model/web"
	conversion "acne-scan-api/internal/pkg/conversion/request"
	"mime/multipart"
	"time"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (pr *ProductRecommendationServiceImpl) Create(productRecommendation web.ProductRecommendationRequest,image *multipart.FileHeader,c *fiber.Ctx) error {
	var err error

	err = pr.Validator.Struct(productRecommendation)
	if err != nil {
		return err
	}

	conv := conversion.ProductRecommendationToModel(productRecommendation)

	imageUrlChan := make(chan string)
	errChan := make(chan string)

	go func() {
		imageUrl, err := pr.BucketUploder.Uploader(c, image)
		if err != nil {
			errChan <- err.Error()
			return
		}

		imageUrlChan <- imageUrl
	}()

	select {
	case imageUrl := <-imageUrlChan:
		conv.Image = imageUrl
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
	conv.CreatedAt = createdAt
	conv.UpdatedAt = createdAt
	conv.RecommendationId = randomID

	err = pr.ProductRecommendationRepository.Create(conv)
	if err != nil {
		return fmt.Errorf("gagal menginput product recommendation: %s",err.Error())
	}

	return nil
}
