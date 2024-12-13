package service

import (
	"acne-scan-api/internal/model/web"
	"fmt"
	"time"
)

func (pr *ProductRecommendationServiceImpl) Update(article *web.ProductRecommendationUpdateRequest, id string) error {
	ifExist, err := pr.ProductRecommendationRepository.GetById(id)
	if err != nil {
		return fmt.Errorf("failed to find product recommendation:%s", err.Error())
	}

	if ifExist == nil {
		return fmt.Errorf("product recommendation not found")
	}

	wib, err := time.LoadLocation("Asia/Jakarta") // WIB (UTC+7)
	if err != nil {
		return fmt.Errorf("error loading WIB location: %s", err.Error())
	}

	now := time.Now().In(wib)
	article.UpdatedAt = now

	err = pr.ProductRecommendationRepository.Update(article, id)
	if err != nil {
		return fmt.Errorf("error when updating : %s", err.Error())
	}

	return nil
}
