package core

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

func (svc *service) CreateDefaultCategories(workspace string) error {
	for _, category := range defaultCategories {
		category.Workspace = workspace
		err := svc.CreateCategory(category)
		if err != nil {
			return &ErrDbOpFailed{err}
		}
	}
	return nil
}

func (svc *service) CreateCategory(category Category) error {
	category.ID = svc.ut.GenerateID()
	category.CreatedAt = svc.ut.GetCurrentTime()

	err := svc.db.CreateItem(category)
	if err != nil {
		return &ErrDbOpFailed{err}
	}
	return nil
}

func (svc *service) ListCategories(workspace string) ([]Category, error) {
	out, err := svc.db.ListItemsByWorkspace(workspace)
	if err != nil {
		return []Category{}, &ErrDbOpFailed{err}
	}
	return out, nil
}

func (svc *service) EditCategory(category Category) error {
	category.UpdatedAt = svc.ut.GetCurrentTime()

	err := svc.db.EditItem(category)
	if err != nil {
		return &ErrDbOpFailed{err}
	}
	return nil
}

func (svc *service) DeleteCategory(category Category) error {
	err := svc.db.DeleteItem(category)
	if err != nil {
		return &ErrDbOpFailed{err}
	}
	return nil
}
