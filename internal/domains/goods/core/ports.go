package core

type ServicePort interface {
	AddGood(Good) error
	GetGood(Good) (Good, error)
	ListGoods() ([]Good, error)
	DeleteGood(Good) error
}

type DatabasePort interface {
	CreateItem(Good) error
	GetItemByID(Good) (Good, error)
	GetAllItems() ([]Good, error)
	DeleteItem(Good) error
}

type ImagePort interface {
	GetImageURL(string) string
}
