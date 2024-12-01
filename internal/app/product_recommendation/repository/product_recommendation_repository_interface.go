package repository

import (
	"acne-scan-api/internal/model/domain"
	"acne-scan-api/internal/model/web"
	"database/sql"
)

type ProductRecommendationRepository interface {
	Create(productRecommendation *domain.ProductRecommendation) error
	GetAll() ([]domain.ProductRecommendation, error)
	Delete(id int)error
	GetById(id int) (*domain.ProductRecommendation,error)
	Update(recommendation *web.ProductRecommendationUpdateRequest, id int) error
}

type ProductRecommendationImpl struct {
	DB *sql.DB
}

func NewProProductRecommendation(db *sql.DB) ProductRecommendationRepository {
	return &ProductRecommendationImpl{
		DB: db,
	}
}
