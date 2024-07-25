package mocks

import "pantori/internal/domains/categories/core"

type Service struct {
	AnalyzeStr      func(string) core.DescribedError
	AnalyzeCategory func(core.Category) core.DescribedError
	Invoked         bool
	Err             core.DescribedError
}

func (svc *Service) CreateDefaultCategories(workspace string) core.DescribedError {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeStr(workspace)
}

func (svc *Service) CreateCategory(category core.Category) core.DescribedError {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeCategory(category)
}

func (svc *Service) ListCategories(workspace string) ([]core.Category, core.DescribedError) {
	svc.Invoked = true
	if svc.Err != nil {
		return []core.Category{}, svc.Err
	}
	return []core.Category{}, svc.AnalyzeStr(workspace)
}

func (svc *Service) EditCategory(category core.Category) core.DescribedError {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeCategory(category)
}

func (svc *Service) DeleteCategory(category core.Category) core.DescribedError {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return svc.AnalyzeCategory(category)
}
