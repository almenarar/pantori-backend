package mocks

import (
	"pantori/internal/domains/categories/core"
)

type DatabaseMock struct {
	ErrAdd       error
	ErrEdit      error
	ErrList      error
	ErrDelete    error
	CreatedItems []core.Category
	Invocation   *string
}

func (db *DatabaseMock) ListItemsByWorkspace(string) ([]core.Category, error) {
	*db.Invocation = *db.Invocation + "-List"
	if db.ErrList != nil {
		return []core.Category{}, db.ErrList
	}
	return []core.Category{
		{ID: "foo"},
		{ID: "bar"},
	}, nil
}

func (db *DatabaseMock) CreateItem(category core.Category) error {
	*db.Invocation = *db.Invocation + "-Add"
	db.CreatedItems = append(db.CreatedItems, category)
	if db.ErrAdd != nil {
		return db.ErrAdd
	}
	return nil
}

func (db *DatabaseMock) EditItem(core.Category) error {
	*db.Invocation = *db.Invocation + "-Edit"
	if db.ErrEdit != nil {
		return db.ErrEdit
	}
	return nil
}

func (db *DatabaseMock) DeleteItem(core.Category) error {
	*db.Invocation = *db.Invocation + "-Delete"
	if db.ErrDelete != nil {
		return db.ErrDelete
	}
	return nil
}
