package domain

import "time"

type ProductRecommendation struct {
	RecommendationId int       `json:"recommendation_id"`
	Image            string    `json:"image"`
	Link             string    `json:"link"`
	Description      string    `json:"description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
