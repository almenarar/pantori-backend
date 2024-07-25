package core

type ServicePort interface {
	AddGood(Good) DescribedError
	EditGood(Good) DescribedError
	GetGood(Good) (Good, DescribedError)
	ListGoods(string) ([]Good, DescribedError)
	DeleteGood(Good) DescribedError
	BuildShoppingList(workspace string) ([]Good, DescribedError)
}

type DatabasePort interface {
	CreateItem(Good) error
	EditItem(Good) error
	GetItemByID(Good) (Good, error)
	GetAllItems(string) ([]Good, error)
	DeleteItem(Good) error
}

type ImagePort interface {
	GetImageURL(string) string
}

type UtilsPort interface {
	GenerateID() string
	GetCurrentTime() string
}
