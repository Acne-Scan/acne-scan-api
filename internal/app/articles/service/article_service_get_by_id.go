package service

import (
	"acne-scan-api/internal/model/domain"
	"fmt"
)

func (articleService *ArticleServiceImpl) GetById(id string) (*domain.Article, error) {

	data, err := articleService.ArticleRepository.GetById(id)
	if err != nil || data == nil {
		return nil, fmt.Errorf("article not found %s", err.Error())
	}

	return data, nil
}
