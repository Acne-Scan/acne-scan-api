package web

type Images map[string]string
type ProductLink map[string]interface{}

type HistoryRequest struct {
	User_id        uint        `json:"user_id" form:"user_id"`
	HistoryId      string      `json:"history_id" form:"history_id"`
	Image          Images      `json:"image" form:"image"`
	UserPicture    string      `json:"user_picture" form:"user_picture"`
	ProductLink    ProductLink `json:"product_links" form:"product_links"`
	Prediction     string      `json:"prediction" form:"prediction"`
	Recommendation string      `json:"recommendation" form:"recommendation"`
}
