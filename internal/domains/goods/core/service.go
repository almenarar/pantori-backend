package core

type service struct {
	db DatabasePort
	im ImagePort
	ut UtilsPort
}

func NewService(db DatabasePort, im ImagePort, ut UtilsPort) *service {
	return &service{
		db: db,
		im: im,
		ut: ut,
	}
}

func (svc *service) AddGood(good Good) error {
	// a fail here is not critical, will only log
	good.ImageURL = svc.im.GetImageURL(good.Name)

	good.ID = svc.ut.GenerateID()
	good.CreatedAt = svc.ut.GetCurrentTime()

	err := svc.db.CreateItem(good)
	if err != nil {
		return &ErrDbOpFailed{err}
	}
	return nil
}

func (svc *service) EditGood(good Good) error {
	good.UpdatedAt = svc.ut.GetCurrentTime()

	err := svc.db.EditItem(good)
	if err != nil {
		return &ErrDbOpFailed{err}
	}
	return nil
}

func (svc *service) GetGood(good Good) (Good, error) {
	result, err := svc.db.GetItemByID(good)
	if err != nil {
		return Good{}, &ErrDbOpFailed{err}
	}
	return result, nil
}

func (svc *service) ListGoods(workspace string) ([]Good, error) {
	result, err := svc.db.GetAllItems(workspace)
	if err != nil {
		return []Good{}, &ErrDbOpFailed{err}
	}
	return result, nil
}

func (svc *service) DeleteGood(good Good) error {
	err := svc.db.DeleteItem(good)
	if err != nil {
		return &ErrDbOpFailed{err}
	}
	return nil
}
