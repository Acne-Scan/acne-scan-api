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

func (history *HistoryServiceImpl) Create(c *fiber.Ctx, request *web.HistoryRequest, historyJson []byte, image *multipart.FileHeader, productLinkJson []byte, userId string) error {
	var err error

	err = history.Validator.Struct(request)
	if err != nil {
		return err
	}

	historyConv := conversion.HistoryRequestToModel(request)

	if image != nil {
		imageUrlChan := make(chan string)
		errChan := make(chan string)

		go func() {
			imageUrl, err := history.BucketUploder.Uploader(c, image)
			if err != nil {
				errChan <- err.Error()
				return
			}

			imageUrlChan <- imageUrl
		}()

		select {
		case imageUrl := <-imageUrlChan:
			historyConv.UserPicture = imageUrl
		case err := <-errChan:
			return fmt.Errorf("error uploading image: %s", err)
		}
	}

	wib, err := time.LoadLocation("Asia/Jakarta") // WIB (UTC+7)
	if err != nil {
		return fmt.Errorf("error loading WIB location: %s", err.Error())
	}

	//generate random id
	randomID := uuid.New().String()

	createdAt := time.Now().In(wib)
	historyConv.CreatedAt = createdAt
	historyConv.UpdatedAt = createdAt
	historyConv.HistoryId = randomID


	err = history.HistoryRepository.Create(historyConv, historyJson, productLinkJson, userId)
	if err != nil {
		return fmt.Errorf("error register %s", err.Error())
	}

	return nil
}
