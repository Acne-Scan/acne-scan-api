package repository

import (
	"acne-scan-api/internal/model/domain"
	"database/sql"
	"time"
)

type ArticleRepository interface {
	Create(article *domain.Article)error
	Update(name,description,image string, id string,updatedAt time.Time)error
	Delete(id string)error
	GetAll()([]domain.Article,error)
	GetById(id string)(*domain.Article,error)
}

type ArticleRepositoryImpl struct{
	DB *sql.DB
}

func NewArticleRepository(db *sql.DB)ArticleRepository{
	return &ArticleRepositoryImpl{DB: db}
}