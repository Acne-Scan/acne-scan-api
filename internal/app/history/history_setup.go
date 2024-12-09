package history

import (
	"acne-scan-api/internal/app/history/handlers"
	"acne-scan-api/internal/app/history/repository"
	"acne-scan-api/internal/app/history/routes"
	"acne-scan-api/internal/app/history/service"
	cloudstorage "acne-scan-api/internal/pkg/cloud_storage"
	"database/sql"

	"github.com/go-playground/validator"
)

func HistorySetup(db *sql.DB, validate *validator.Validate,bs cloudstorage.StorageBucketUploader) routes.HistoryRoutes {
	HistoryRepository := repository.NewHistoryRepository(db)
	HistoryService := service.NewHistoryService(HistoryRepository, validate,bs)
	HistoryHandlers := handlers.NewHistoryHandlers(HistoryService)
	HistoryRoutes := routes.NewHistoryRoutes(HistoryHandlers)

	return HistoryRoutes
}