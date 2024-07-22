package core

//-------------------------------------------------------------------------------
// Principal
//-------------------------------------------------------------------------------

type Good struct {
	ID         string   `gorm:"type:char(36);primaryKey" json:"ID"`
	Name       string   `json:"Name"`
	Categories []string `json:"Categories"`
	Quantity   string   `json:"Quantity"`
	ImageURL   string   `json:"ImageURL"`
	Workspace  string   `json:"Workspace"`
	Expire     string   `json:"Expire"`
	OpenExpire string   `json:"OpenExpire"`
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
	Quantity   string   `json:"Quantity"`
	Expire     string   `json:"Expire" binding:"required"`
	BuyDate    string   `json:"BuyDate" binding:"required"`
}

type PatchGood struct {
	ID         string   `json:"ID" binding:"required"`
	Name       string   `json:"Name" binding:"required"`
	Categories []string `json:"Categories" binding:"required"`
	Quantity   string   `json:"Quantity"`
	ImageURL   string   `json:"ImageURL"`
	Expire     string   `json:"Expire" binding:"required"`
	OpenExpire string   `json:"OpenExpire"`
	BuyDate    string   `json:"BuyDate" binding:"required"`
	CreatedAt  string   `json:"CreatedAt" binding:"required"`
}

type GetGood struct {
	ID string `json:"ID" binding:"required"`
}

type DeleteGood struct {
	ID string `json:"ID" binding:"required"`
}
