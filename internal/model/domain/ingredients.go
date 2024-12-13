package domain

import "time"

type Ingredients struct {
	Ingredientsid string    `gorm:"type:int;primarykey" json:"ingredients_id"`
	HistoryId     int    `gorm:"type:int;" json:"history_id"`
	Ingredients   string `gorm:"type:varchar(255)" json:"ingredients"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}