package service

import (
	"acne-scan-api/internal/app/product_recommendation/repository"
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
	cloudstorage "acne-scan-api/internal/pkg/cloud_storage"
	"mime/multipart"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ProductRecommendationService interface {
	Create(pr web.ProductRecommendationRequest,image *multipart.FileHeader,c *fiber.Ctx) error
	GetAll() ([]domain.ProductRecommendation, error)
	Delete(id string) error
	GetById(id string) (*domain.ProductRecommendation, error)
	Update(article *web.ProductRecommendationUpdateRequest, id string) error
}

type ProductRecommendationServiceImpl struct {
	ProductRecommendationRepository repository.ProductRecommendationRepository
	Validator                       *validator.Validate
	BucketUploder                   cloudstorage.StorageBucketUploader
}

func NewProductRecommendationService(pr repository.ProductRecommendationRepository, validate *validator.Validate, bu cloudstorage.StorageBucketUploader) ProductRecommendationService {
	return &ProductRecommendationServiceImpl{
		ProductRecommendationRepository: pr,
		Validator:                       validate, BucketUploder: bu,
	}
}
