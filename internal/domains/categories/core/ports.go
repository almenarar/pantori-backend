package core

type ServicePort interface {
	CreateDefaultCategories(workspace string) error
	CreateCategory(Category) error
	ListCategories(workspace string) ([]Category, error)
	EditCategory(Category) error
	DeleteCategory(Category) error
}

type DatabasePort interface {
	CreateItem(Category) error
	EditItem(Category) error
	ListItemsByWorkspace(workspace string) ([]Category, error)
	DeleteItem(Category) error
}

type UtilsPort interface {
	GenerateID() string
	GetCurrentTime() string
}
