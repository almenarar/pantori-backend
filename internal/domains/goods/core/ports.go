package core

type ServicePort interface {
	AddGood(Good) error
	EditGood(Good) error
	GetGood(Good) (Good, error)
	ListGoods(string) ([]Good, error)
	DeleteGood(Good) error
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
