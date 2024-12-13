package repository

import (
	"acne-scan-api/internal/model/web"
	"fmt"
)

func (pr *ProductRecommendationImpl) Update(recommendation *web.ProductRecommendationUpdateRequest, id string) error {
	_, err := pr.DB.Exec("update pruduct_recommendation set image=?, link=?, description=? where recommendation_id=?", recommendation.Image, recommendation.Link, recommendation.Description,id)

	if err != nil {
		fmt.Println("query gagal: ", err.Error())
		return err
	}

	return nil
}
