package repository

import (
	"acne-scan-api/internal/model/domain"
)

func (history *HistoryRepositoryImpl) Create(domainHistory *domain.History, historyJson []byte, productLinkJson []byte, userId int) error {
	stmt, err := history.DB.Prepare("insert into history (history_id,user_id,image,product_link,user_picture,prediction,recommendation,created_at,updated_at) values (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		domainHistory.HistoryId,
		userId,
		string(historyJson),
		string(productLinkJson),
		domainHistory.UserPicture,
		domainHistory.Prediction,
		domainHistory.Recommendation,
		domainHistory.CreatedAt,
		domainHistory.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}
