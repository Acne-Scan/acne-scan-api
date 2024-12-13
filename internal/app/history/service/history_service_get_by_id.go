package service

import (
	"acne-scan-api/internal/model/domain"
	"fmt"
)

func (history *HistoryServiceImpl) GetById(id string) (*domain.History,error ){
	data, err := history.HistoryRepository.GetById(id)
	if err != nil || data == nil {
		return nil, fmt.Errorf("no result %s", err.Error())
	}

	return data, nil
}