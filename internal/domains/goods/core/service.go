package core

type service struct {
	db DatabasePort
	im ImagePort
}

func NewService(db DatabasePort, im ImagePort) *service {
	return &service{
		db: db,
		im: im,
	}
}

func (svc *service) AddGood(good Good) error {
	// a fail here is not critical, will only log
	good.ImageURL = svc.im.GetImageURL(good.Name)

	err := svc.db.CreateItem(good)
	if err != nil {
		return &ErrDbOpFailed{err}
	}
	return nil
}

func (svc *service) EditGood(good Good) error {
	err := svc.db.EditItem(good)
	if err != nil {
		return &ErrDbOpFailed{err}
	}
	return nil
}

func (svc *service) GetGood(id string) (Good, error) {
	result, err := svc.db.GetItemByID(id)
	if err != nil {
		return Good{}, &ErrDbOpFailed{err}
	}
	return result, nil
}

func (svc *service) ListGoods() ([]Good, error) {
	result, err := svc.db.GetAllItems()
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
