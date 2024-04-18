package mocks

import "pantori/internal/domains/categories/core"

type Service struct {
	AnalyzeStr      func(string) error
	AnalyzeCategory func(core.Category) error
	Invoked         bool
	Err             error
}

func (svc *Service) CreateDefaultCategories(workspace string) error {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeStr(workspace)
}

func (svc *Service) CreateCategory(category core.Category) error {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeCategory(category)
}

func (svc *Service) ListCategories(workspace string) ([]core.Category, error) {
	svc.Invoked = true
	if svc.Err != nil {
		return []core.Category{}, svc.Err
	}
	return []core.Category{}, svc.AnalyzeStr(workspace)
}

func (svc *Service) EditCategory(category core.Category) error {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeCategory(category)
}

func (svc *Service) DeleteCategory(category core.Category) error {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeCategory(category)
}
