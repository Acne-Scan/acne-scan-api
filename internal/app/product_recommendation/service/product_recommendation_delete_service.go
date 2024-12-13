package service

import "fmt"

func (pr *ProductRecommendationServiceImpl) Delete(id string) error {

	ifExist, _ := pr.ProductRecommendationRepository.GetById(id)
	if ifExist == nil {
		return fmt.Errorf("product recommendation not found")
	}

	err := pr.ProductRecommendationRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}