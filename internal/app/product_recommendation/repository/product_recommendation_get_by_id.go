package repository

import (
	"acne-scan-api/internal/model/domain"
)

func (pr *ProductRecommendationImpl) GetById(id string) (*domain.ProductRecommendation, error) {

	result := domain.ProductRecommendation{}

	err := pr.DB.QueryRow("select recommendation_id,image,link,description,created_at,updated_at from pruduct_recommendation where recommendation_id=?", id).Scan(
		&result.RecommendationId,
		&result.Image,
		&result.Link,
		&result.Description,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
