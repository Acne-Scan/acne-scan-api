package repository

import (
	"acne-scan-api/internal/model/domain"
	"database/sql"
)

type HistoryRepository interface {
	Create(domainHistory *domain.History,historyJson []byte,productLinkJson []byte,userid string) error
	GetById(id string) (*domain.History, error)
	GetAll(id string) ([]*domain.History, error) 
}

type HistoryRepositoryImpl struct {
	DB *sql.DB
}

func NewHistoryRepository(db *sql.DB) HistoryRepository {
	return &HistoryRepositoryImpl{DB: db}
}
