package service

import (
	"acne-scan-api/internal/app/articles/repository"
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
	cloudstorage "acne-scan-api/internal/pkg/cloud_storage"
	"mime/multipart"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type ArticleService interface {
	Create(request web.ArticleCreateRequest, image *multipart.FileHeader,c *fiber.Ctx) error 
	GetAll()([]domain.Article,error)
	GetById(id string)(*domain.Article,error)
	Delete(id string)(error)
	Update(name,description,image string, id string)(error)
}

type ArticleServiceImpl struct{
	ArticleRepository repository.ArticleRepository
	Validator *validator.Validate
	BucketUploder cloudstorage.StorageBucketUploader
}

func NewArticleService(ar repository.ArticleRepository,validate *validator.Validate,bs cloudstorage.StorageBucketUploader) ArticleService{
	return &ArticleServiceImpl{ArticleRepository: ar,Validator: validate,BucketUploder: bs}
}

