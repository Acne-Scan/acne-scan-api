package web

import "time"

type ArticleCreateRequest struct {
	ArticleId   string       `form:"article_id"`
	Name        string    `form:"name"`
	Image       string    `form:"image"`
	Description string    `form:"description" validate:"required"`
	CreatedAt   time.Time `form:"created_at"`
	UpdatedAt   time.Time `form:"updated_at"`
}

type ArticleUpdateRequest struct {
	Image       string    `form:"image"`
	Description string    `form:"description"`
	UpdatedAt   time.Time `form:"updated_at"`
}
