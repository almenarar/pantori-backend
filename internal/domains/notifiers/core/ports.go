package core

type ServicePort interface {
	NotifyExpiredGoods()
}

type GoodsPort interface {
	GetGoodsFromWorkspace(workspace string) ([]Good, error)
}

type UsersPort interface {
	ListAllUsersByWorkspace() (map[string][]User, error)
}

type EmailPort interface {
	SendEmail(user User, report Report) error
}
