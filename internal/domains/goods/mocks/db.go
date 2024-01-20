package mocks

import (
	"pantori/internal/domains/goods/core"
)

type DatabaseMock struct {
	ErrGet     error
	ErrAdd     error
	ErrList    error
	ErrDelete  error
	Invocation *string
}

func (db *DatabaseMock) GetItemByID(core.Good) (core.Good, error) {
	*db.Invocation = *db.Invocation + "-Get"
	if db.ErrGet != nil {
		return core.Good{}, db.ErrGet
	}
	return core.Good{ID: "foo"}, nil
}

func (db *DatabaseMock) GetAllItems() ([]core.Good, error) {
	*db.Invocation = *db.Invocation + "-List"
	if db.ErrList != nil {
		return []core.Good{}, db.ErrList
	}
	return []core.Good{
		{ID: "foo"},
		{ID: "bar"},
	}, nil
}

func (db *DatabaseMock) CreateItem(core.Good) error {
	*db.Invocation = *db.Invocation + "-Add"
	if db.ErrAdd != nil {
		return db.ErrAdd
	}
	return nil
}

func (db *DatabaseMock) DeleteItem(core.Good) error {
	*db.Invocation = *db.Invocation + "-Delete"
	if db.ErrDelete != nil {
		return db.ErrDelete
	}
	return nil
}
