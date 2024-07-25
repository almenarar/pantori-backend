package mocks

import "pantori/internal/auth/core"

type Service struct {
	CustomFunc func(core.User) core.DescribedError
	Invoked    bool
	Err        core.DescribedError
}

func (svc *Service) Authenticate(user core.User) (string, core.DescribedError) {
	svc.Invoked = true
	if svc.Err != nil {
		return "", svc.Err
	}
	return "token", svc.CustomFunc(user)
}

func (svc *Service) CreateUser(core.User) core.DescribedError {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return nil
}

func (svc *Service) DeleteUser(core.User) core.DescribedError {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return nil
}

func (svc *Service) ListUsers() ([]core.User, core.DescribedError) {
	svc.Invoked = true
	if svc.Err != nil {
		return []core.User{}, svc.Err
	}
	return []core.User{{Username: "john", Workspace: "wkp1"}}, nil
}
