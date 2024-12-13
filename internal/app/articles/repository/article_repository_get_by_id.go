package repository

import (
	"acne-scan-api/internal/model/domain"
)

func (articleRepository *ArticleRepositoryImpl) GetById(id string) (*domain.Article, error) {

	result := domain.Article{}

	err := articleRepository.DB.QueryRow("select article_id,name, image, description,created_at,updated_at from article where article_id=?", id).Scan(
		&result.ArticleId,
		&result.Name,
		&result.Image,
		&result.Description,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
