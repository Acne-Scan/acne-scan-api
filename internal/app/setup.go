package app

import (
	"acne-scan-api/internal/app/articles"
	"acne-scan-api/internal/app/auth"
	"acne-scan-api/internal/app/history"
	productrecommendation "acne-scan-api/internal/app/product_recommendation"
	cloudstorage "acne-scan-api/internal/pkg/cloud_storage"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func InitApp(db *sql.DB, validate *validator.Validate, app *fiber.App, bs cloudstorage.StorageBucketUploader) {
	ArticleRoutes := articles.ArticleSetup(db, validate, bs)
	ProductRecommendationRoutes := productrecommendation.ProductRecommendationSetup(db, validate, bs)
	authRoutes := auth.AuthSetup(db, validate)
	historyRoutes := history.HistorySetup(db, validate, bs)

	ProductRecommendationRoutes.ProductRecommendationRoutes(app)
	ArticleRoutes.ArticleRoutes(app)
	authRoutes.AuthRoutes(app)
	historyRoutes.HistoryRoutes(app)
}
