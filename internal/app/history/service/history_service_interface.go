package service

import (
	"acne-scan-api/internal/app/history/repository"
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
	cloudstorage "acne-scan-api/internal/pkg/cloud_storage"
	"mime/multipart"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type HistoryService interface {
	Create(c *fiber.Ctx, request *web.HistoryRequest, historyJson []byte,image *multipart.FileHeader, productLinkJson []byte, userId int) error
	GetById(id int) (*domain.History, error)
	GetAll(id int) ([]*domain.History, error)
}

type HistoryServiceImpl struct {
	HistoryRepository repository.HistoryRepository
	Validator         *validator.Validate
	BucketUploder     cloudstorage.StorageBucketUploader
}

func NewHistoryService(HistoryRepository repository.HistoryRepository, validate *validator.Validate,bs cloudstorage.StorageBucketUploader) HistoryService {
	return &HistoryServiceImpl{HistoryRepository: HistoryRepository, Validator: validate,BucketUploder: bs}
}
