package core

import "context"

type ServicePort interface {
	NotifyExpiredGoods(ctx context.Context) error
}

type GoodsPort interface {
	GetGoodsFromWorkspace(workspace string) ([]Good, error)
}

type UsersPort interface {
	ListAllUsersByWorkspace() (map[string][]User, error)
}

type EmailPort interface {
	SendEmail(user User, expireToday, expireSoon, expired []Good) error
}
