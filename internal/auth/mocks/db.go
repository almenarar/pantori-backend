package mocks

import (
	core "pantori/internal/auth/core"
)

type DatabaseMock struct {
	Err        error
	UserExists bool
	Invocation *string
}

func (db *DatabaseMock) GetUser(string) (core.User, error) {
	*db.Invocation = *db.Invocation + "-Get"
	if db.Err != nil {
		return core.User{}, db.Err
	}
	if db.UserExists {
		return core.User{ActualPassword: "foo"}, nil
	}
	return core.User{}, nil
}

func (db *DatabaseMock) ListUsers() ([]core.User, error) {
	*db.Invocation = *db.Invocation + "-List"
	if db.Err != nil {
		return []core.User{}, db.Err
	}
	return []core.User{{Username: "john"}}, nil
}

func (db *DatabaseMock) CreateUser(user core.User) error {
	*db.Invocation = *db.Invocation + "-Create"
	if db.Err != nil {
		return db.Err
	}
	return nil
}

func (db *DatabaseMock) DeleteUser(user core.User) error {
	*db.Invocation = *db.Invocation + "-Delete"
	if db.Err != nil {
		return db.Err
	}
	return nil
}
