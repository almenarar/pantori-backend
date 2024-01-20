package core

import "time"

//-------------------------------------------------------------------------------
// Principal
//-------------------------------------------------------------------------------

type Good struct {
	ID        string    `gorm:"type:char(36);primaryKey" json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	ImageURL  string    `json:"image_url"`
	Workspace string    `json:"workspace"`
	Expire    time.Time `json:"expire"`
	BuyDate   time.Time `json:"buy_date"`
	CreatedAt time.Time `gorm:"default:current_timestamp(3)" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:current_timestamp(3)" json:"updated_at"`
}

//-------------------------------------------------------------------------------
// Swagger only
//-------------------------------------------------------------------------------

type PostGood struct {
	Name      string `json:"name" binding:"required"`
	Category  string `json:"category" binding:"required"`
	Workspace string `json:"workspace" binding:"required"`
	Expire    string `json:"expire" binding:"required"`
	BuyDate   string `json:"buy_date" binding:"required"`
}

type GetGood struct {
	ID string `json:"id" binding:"required"`
}

type DeleteGood struct {
	ID string `json:"id"`
}
