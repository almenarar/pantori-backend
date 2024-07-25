package handlers

import (
	"pantori/internal/domains/goods/core"

	notifiers "pantori/internal/domains/notifiers/core"
)

type Internal struct {
	svc core.ServicePort
}

func NewInternal(svc core.ServicePort) *Internal {
	return &Internal{svc: svc}
}

func (int *Internal) GetGoodsFromWorkspace(workspace string) ([]notifiers.Good, error) {
	var output []notifiers.Good

	goods, err := int.svc.ListGoods(workspace)
	if err != nil {
		return output, err
	}

	for _, good := range goods {
		if good.OpenExpire != "" {
			good.Expire = good.OpenExpire
		}
		output = append(output, notifiers.Good{
			Name:      good.Name,
			Workspace: good.Workspace,
			Quantity:  good.Quantity,
			Expire:    good.Expire,
		})
	}

	return output, nil
}
