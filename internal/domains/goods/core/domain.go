package core

//-------------------------------------------------------------------------------
// Principal
//-------------------------------------------------------------------------------

type Good struct {
	ID         string   `gorm:"type:char(36);primaryKey" json:"ID"`
	Name       string   `json:"Name"`
	Categories []string `json:"Categories"`
	ImageURL   string   `json:"ImageURL"`
	Workspace  string   `json:"Workspace"`
	Expire     string   `json:"Expire"`
	BuyDate    string   `json:"BuyDate"`
	CreatedAt  string   `gorm:"default:current_timestamp(3)" json:"CreatedAt"`
	UpdatedAt  string   `gorm:"default:current_timestamp(3)" json:"UpdatedAt"`
}

//-------------------------------------------------------------------------------
// Swagger only
//-------------------------------------------------------------------------------

type PostGood struct {
	Name       string   `json:"Name" binding:"required"`
	Categories []string `json:"Categories"`
	Expire     string   `json:"Expire" binding:"required"`
	BuyDate    string   `json:"BuyDate" binding:"required"`
}

type PatchGood struct {
	ID         string   `json:"ID" binding:"required"`
	Name       string   `json:"Name" binding:"required"`
	Categories []string `json:"Categories" binding:"required"`
	ImageURL   string   `json:"ImageURL" binding:"required"`
	Expire     string   `json:"Expire" binding:"required"`
	BuyDate    string   `json:"BuyDate" binding:"required"`
	CreatedAt  string   `json:"CreatedAt" binding:"required"`
}

type GetGood struct {
	ID string `json:"ID" binding:"required"`
}

type DeleteGood struct {
	ID string `json:"ID" binding:"required"`
}
