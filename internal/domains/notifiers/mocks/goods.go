package mocks

import (
	"pantori/internal/domains/notifiers/core"
	"time"
)

type GoodsMocks struct {
	ErrGet     error
	HaveGoods  bool
	Invocation *string
}

func (gm *GoodsMocks) GetGoodsFromWorkspace(workspace string) ([]core.Good, error) {
	*gm.Invocation = *gm.Invocation + "-GetGoods"

	goods := []core.Good{}
	if gm.ErrGet != nil {
		return goods, gm.ErrGet
	}

	goods = append(goods, core.Good{Name: "beans", Expire: "30/11/2122"})
	if gm.HaveGoods {
		goods = append(goods, core.Good{Name: "grape", Expire: time.Now().Format("02/01/2006")})
	}

	return goods, nil
}
