package web

import "time"

type ProductRecommendationRequest struct {
	RecommendationId string       `form:"recommendation_id" json:"recommendation_id" validate:"required"`
	Image            string    `form:"image" json:"image"`
	Link             string    `form:"link" json:"link" validate:"required"`
	Description      string    `form:"description" json:"description" validate:"required"`
	CreatedAt        time.Time `form:"created_at" json:"created_at"`
	UpdatedAt        time.Time `form:"updated_at" json:"updated_at"`
}

type ProductRecommendationUpdateRequest struct {
	Image       string    `form:"image" validate:"required"`
	Link        string    `form:"link" json:"link" validate:"required"`
	Description string    `form:"description" json:"description" validate:"required"`
	UpdatedAt   time.Time `form:"updated_at" json:"updated_at"`
}
