package core

import "github.com/rs/zerolog/log"

type service struct {
	db DatabasePort
	ut UtilsPort
}

func NewService(db DatabasePort, ut UtilsPort) *service {
	return &service{
		db: db,
		ut: ut,
	}
}

func (svc *service) CreateDefaultCategories(workspace string) DescribedError {
	for _, category := range defaultCategories {
		category.Workspace = workspace
		err := svc.CreateCategory(category)
		if err != nil {
			return &ErrDBCreateFailed{err}
		}
	}
	return nil
}

func (svc *service) CreateCategory(category Category) DescribedError {
	category.ID = svc.ut.GenerateID()
	category.CreatedAt = svc.ut.GetCurrentTime()

	err := svc.db.CreateItem(category)
	if err != nil {
		log.Error().Err(&ErrDBCreateFailed{err}).Msg("")
		return &ErrDBCreateFailed{err}
	}
	return nil
}

func (svc *service) ListCategories(workspace string) ([]Category, DescribedError) {
	out, err := svc.db.ListItemsByWorkspace(workspace)
	if err != nil {
		log.Error().Err(&ErrDBListFailed{err}).Msg("")
		return []Category{}, &ErrDBListFailed{err}
	}
	return out, nil
}

func (svc *service) EditCategory(category Category) DescribedError {
	category.UpdatedAt = svc.ut.GetCurrentTime()

	err := svc.db.EditItem(category)
	if err != nil {
		log.Error().Err(&ErrDBEditFailed{err}).Msg("")
		return &ErrDBEditFailed{err}
	}
	return nil
}

func (svc *service) DeleteCategory(category Category) DescribedError {
	err := svc.db.DeleteItem(category)
	if err != nil {
		log.Error().Err(&ErrDBDeleteFailed{err}).Msg("")
		return &ErrDBDeleteFailed{err}
	}
	return nil
}
