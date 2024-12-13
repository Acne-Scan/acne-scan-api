package repository

import "fmt"

func (pr *ProductRecommendationImpl) Delete(id string) error {
	result, err := pr.DB.Exec("delete from pruduct_recommendation where recommendation_id=?", id)
	if err != nil {
		return err
	}

	rowAffected, _ := result.RowsAffected()

	if rowAffected < 1 {
		return fmt.Errorf("not row affected or id not found")
	}

	return nil
}