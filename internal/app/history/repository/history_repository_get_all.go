package repository

import (
	"acne-scan-api/internal/model/domain"
	"encoding/json"
	"fmt"
)

func (history *HistoryRepositoryImpl) GetAll(id string) ([]*domain.History, error) {

	result := []*domain.History{}

	rows, err := history.DB.Query("select history.history_id, history.user_id, history.user_picture ,users.username, history.image, history.product_link, history.prediction, history.recommendation, history.created_at, history.updated_at from history left join users on users.user_id=history.user_id where users.user_id=?", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var imageData []byte
		var ProductLinkData []byte
		each := domain.History{}
		err = rows.Scan(
			&each.HistoryId,
			&each.User_id,
			&each.UserPicture,
			&each.Name,
			&imageData,
			&ProductLinkData,
			&each.Prediction,
			&each.Recommendation,
			&each.CreatedAt,
			&each.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(imageData, &each.Image); err != nil {
			return nil, fmt.Errorf("failed to parse image JSON: %v", err)
		}

		if err := json.Unmarshal(ProductLinkData, &each.ProductLink); err != nil {
			return nil, fmt.Errorf("failed to parse product link JSON: %v", err)
		}

		result = append(result, &each)
	}

	return result, nil
}
