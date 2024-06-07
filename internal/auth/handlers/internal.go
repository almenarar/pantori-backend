package handlers

import (
	"pantori/internal/auth/core"

	notifiers "pantori/internal/domains/notifiers/core"
)

type Internal struct {
	svc core.ServicePort
}

func NewInternal(svc core.ServicePort) *Internal {
	return &Internal{svc: svc}
}

func (int *Internal) ListAllUsersByWorkspace() (map[string][]notifiers.User, error) {
	output := make(map[string][]notifiers.User)

	users, err := int.svc.ListUsers()
	if err != nil {
		return output, err
	}

	for _, user := range users {
		output[user.Workspace] = append(output[user.Workspace], notifiers.User{
			Name:      user.Username,
			Email:     user.Email,
			Workspace: user.Workspace,
		})
	}

	return output, nil
}
