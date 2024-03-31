package core

type ServicePort interface {
	AddGood(Good) error
	EditGood(Good) error
	GetGood(string) (Good, error)
	ListGoods() ([]Good, error)
	DeleteGood(Good) error
}

type DatabasePort interface {
	CreateItem(Good) error
	EditItem(Good) error
	GetItemByID(string) (Good, error)
	GetAllItems() ([]Good, error)
	DeleteItem(Good) error
}

type ImagePort interface {
	GetImageURL(string) string
}
