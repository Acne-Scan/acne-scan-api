package service

import (
	"acne-scan-api/internal/model/domain"
	"fmt"
)

func (history *HistoryServiceImpl) GetAll(id int) ([]*domain.History, error) {
	data, err := history.HistoryRepository.GetAll(id)
	if err != nil || data == nil {
		return nil, fmt.Errorf("history not found %s", err.Error())
	}

	return data, nil
}
