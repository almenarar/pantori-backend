package mocks

import "pantori/internal/domains/goods/core"

type Service struct {
	AnalyzeStr  func(string) core.DescribedError
	AnalyzeGood func(core.Good) core.DescribedError
	Invoked     bool
	Err         core.DescribedError
}

func (svc *Service) AddGood(good core.Good) core.DescribedError {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeGood(good)
}

func (svc *Service) EditGood(good core.Good) core.DescribedError {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeGood(good)
}

func (svc *Service) GetGood(good core.Good) (core.Good, core.DescribedError) {
	svc.Invoked = true
	if svc.Err != nil {
		return core.Good{}, svc.Err
	}
	return core.Good{}, svc.AnalyzeGood(good)
}

func (svc *Service) ListGoods(workspace string) ([]core.Good, core.DescribedError) {
	svc.Invoked = true
	if svc.Err != nil {
		return []core.Good{}, svc.Err
	}
	return []core.Good{{Name: "carrot"}, {Name: "grapes", OpenExpire: "2000"}}, svc.AnalyzeStr(workspace)
}

func (svc *Service) DeleteGood(good core.Good) core.DescribedError {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeGood(good)
}

func (svc *Service) BuildShoppingList(workspace string) ([]core.Good, core.DescribedError) {
	svc.Invoked = true
	if svc.Err != nil {
		return []core.Good{}, svc.Err
	}
	return []core.Good{}, svc.AnalyzeStr(workspace)
}
