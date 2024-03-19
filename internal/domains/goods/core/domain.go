package core

//-------------------------------------------------------------------------------
// Principal
//-------------------------------------------------------------------------------

type Good struct {
	ID        string `gorm:"type:char(36);primaryKey" json:"ID"`
	Name      string `json:"Name"`
	Category  string `json:"Category"`
	ImageURL  string `json:"ImageURL"`
	Workspace string `json:"Workspace"`
	Expire    string `json:"Expire"`
	BuyDate   string `json:"BuyDate"`
	CreatedAt string `gorm:"default:current_timestamp(3)" json:"CreatedAt"`
	UpdatedAt string `gorm:"default:current_timestamp(3)" json:"UpdatedAt"`
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

type PatchGood struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	ImageURL  string `json:"image_url"`
	Workspace string `json:"workspace"`
	Expire    string `json:"expire"`
	BuyDate   string `json:"buy_date"`
	CreatedAt string `json:"created_at"`
}

type GetGood struct {
	ID string `json:"id" binding:"required"`
}

type DeleteGood struct {
	ID string `json:"id"`
}
