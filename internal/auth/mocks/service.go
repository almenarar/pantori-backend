package mocks

import "pantori/internal/auth/core"

type Service struct {
	CustomFunc func(core.User) error
	Invoked    bool
	Err        error
}

func (svc *Service) Authenticate(user core.User) (string, error) {
	svc.Invoked = true
	if svc.Err != nil {
		return "", svc.Err
	}
	return "token", svc.CustomFunc(user)
}

func (svc *Service) CreateUser(core.User) error {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return nil
}

func (svc *Service) DeleteUser(core.User) error {
	svc.Invoked = true
	if svc.Err != nil {
		return svc.Err
	}
	return nil
}

func (svc *Service) ListUsers() ([]core.User, error) {
	svc.Invoked = true
	if svc.Err != nil {
		return []core.User{}, svc.Err
	}
	return []core.User{{Username: "john"}}, nil
}
