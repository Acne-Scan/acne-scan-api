package service

import (
	"acne-scan-api/internal/model/domain"
	"fmt"
)

func (pr *ProductRecommendationServiceImpl) GetById(id string) (*domain.ProductRecommendation, error) {
	data, err := pr.ProductRecommendationRepository.GetById(id)
	if err != nil || data == nil {
		return nil, fmt.Errorf("product recommendation not found %s", err.Error())
	}

	return data, nil
}
