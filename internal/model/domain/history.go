package domain

import "time"

type Images map[string]string
type ProductLink map[string]interface{}

type History struct {
	HistoryId      string         `json:"history_id"`
	User_id        uint        `json:"user_id"`
	Name           string      `json:"name"`
	UserPicture    string      `json:"user_picture"`
	ProductLink    ProductLink `json:"product_links"`
	Image          Images      `json:"image" form:"image"`
	Prediction     string      `json:"prediction"`
	Recommendation string      `json:"recommendation"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
