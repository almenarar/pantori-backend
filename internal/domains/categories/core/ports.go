package core

type ServicePort interface {
	CreateDefaultCategories(workspace string) DescribedError
	CreateCategory(Category) DescribedError
	ListCategories(workspace string) ([]Category, DescribedError)
	EditCategory(Category) DescribedError
	DeleteCategory(Category) DescribedError
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
