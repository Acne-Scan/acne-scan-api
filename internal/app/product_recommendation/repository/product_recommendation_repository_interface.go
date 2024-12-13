package repository

import (
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
	"database/sql"
)

type ProductRecommendationRepository interface {
	Create(productRecommendation *domain.ProductRecommendation) error
	GetAll() ([]domain.ProductRecommendation, error)
	Delete(id string)error
	GetById(id string) (*domain.ProductRecommendation,error)
	Update(recommendation *web.ProductRecommendationUpdateRequest, id string) error
}

type ProductRecommendationImpl struct {
	DB *sql.DB
}

func NewProProductRecommendation(db *sql.DB) ProductRecommendationRepository {
	return &ProductRecommendationImpl{
		DB: db,
	}
}
