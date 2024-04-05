package core

//-------------------------------------------------------------------------------
// Principal
//-------------------------------------------------------------------------------

type Category struct {
	ID        string `gorm:"type:char(36);primaryKey" json:"ID"`
	Name      string `json:"Name"`
	Color     string `json:"Color"`
	Workspace string `json:"Workspace"`
	CreatedAt string `gorm:"default:current_timestamp(3)" json:"CreatedAt"`
	UpdatedAt string `gorm:"default:current_timestamp(3)" json:"UpdatedAt"`
}

//-------------------------------------------------------------------------------
// Swagger only
//-------------------------------------------------------------------------------

type PostCategory struct {
	Name      string `json:"Name" binding:"required"`
	Color     string `json:"Color" binding:"required"`
	Workspace string `json:"Workspace" binding:"required"`
}

type PatchCategory struct {
	ID        string `json:"ID" binding:"required"`
	Name      string `json:"Name" binding:"required"`
	Color     string `json:"Color" binding:"required"`
	Workspace string `json:"Workspace" binding:"required"`
}

type PostDefaultCategories struct {
	Workspace string `json:"Workspace" binding:"required"`
}

type DeleteCategory struct {
	ID        string `json:"ID" binding:"required"`
	Workspace string `json:"Workspace" binding:"required"`
}
