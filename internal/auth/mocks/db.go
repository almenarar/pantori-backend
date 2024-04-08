package mocks

import (
	core "pantori/internal/auth/core"
)

type DatabaseMock struct {
	ErrGet     error
	UserExists bool
	Invocation *string
}

func (db *DatabaseMock) GetUser(core.User) (core.User, error) {
	*db.Invocation = *db.Invocation + "-Get"
	if db.ErrGet != nil {
		return core.User{}, db.ErrGet
	}
	if db.UserExists {
		return core.User{ActualPassword: "foo"}, nil
	}
	return core.User{}, nil
}
