package service

import (
	"acne-scan-api/internal/model/web"
	conversion "acne-scan-api/internal/pkg/conversion/request"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (history *HistoryServiceImpl) Create(c *fiber.Ctx, request *web.HistoryRequest, historyJson []byte, image *multipart.FileHeader, productLinkJson []byte, userId int) error {
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

	createdAt := time.Now().In(wib)
	historyConv.CreatedAt = createdAt
	historyConv.UpdatedAt = createdAt

	err = history.HistoryRepository.Create(historyConv, historyJson, productLinkJson, userId)
	if err != nil {
		return fmt.Errorf("error register %s", err.Error())
	}

	return nil
}
