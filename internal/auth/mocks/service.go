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
