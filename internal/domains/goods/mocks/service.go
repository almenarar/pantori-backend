package mocks

import "pantori/internal/domains/goods/core"

type Service struct {
	AnalyzeStr  func(string) error
	AnalyzeGood func(core.Good) error
	Invoked     bool
	Err         error
}

func (svc *Service) AddGood(good core.Good) error {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeGood(good)
}

func (svc *Service) EditGood(good core.Good) error {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeGood(good)
}

func (svc *Service) GetGood(good core.Good) (core.Good, error) {
	svc.Invoked = true
	if svc.Err != nil {
		return core.Good{}, svc.Err
	}
	return core.Good{}, svc.AnalyzeGood(good)
}

func (svc *Service) ListGoods(workspace string) ([]core.Good, error) {
	svc.Invoked = true
	if svc.Err != nil {
		return []core.Good{}, svc.Err
	}
	return []core.Good{}, svc.AnalyzeStr(workspace)
}

func (svc *Service) DeleteGood(good core.Good) error {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeGood(good)
}

func (svc *Service) BuildShoppingList(workspace string) ([]core.Good, error) {
	svc.Invoked = true
	if svc.Err != nil {
		return []core.Good{}, svc.Err
	}
	return []core.Good{}, svc.AnalyzeStr(workspace)
}
