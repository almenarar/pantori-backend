package core

import (
	"time"

	"github.com/rs/zerolog/log"
)

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

func (svc *service) AddGood(good Good) DescribedError {
	// a fail here is not critical, will only log
	good.ImageURL = svc.im.GetImageURL(good.Name)

	good.ID = svc.ut.GenerateID()
	good.CreatedAt = svc.ut.GetCurrentTime()

	err := svc.db.CreateItem(good)
	if err != nil {
		log.Error().Err(&ErrDBCreateFailed{err}).Msg("")
		return &ErrDBCreateFailed{err}
	}
	return nil
}

func (svc *service) EditGood(good Good) DescribedError {
	good.UpdatedAt = svc.ut.GetCurrentTime()

	err := svc.db.EditItem(good)
	if err != nil {
		log.Error().Err(&ErrDBEditFailed{err}).Msg("")
		return &ErrDBEditFailed{err}
	}
	return nil
}

func (svc *service) GetGood(good Good) (Good, DescribedError) {
	result, err := svc.db.GetItemByID(good)
	if err != nil {
		log.Error().Err(&ErrDBGetFailed{err}).Msg("")
		return Good{}, &ErrDBGetFailed{err}
	}
	return result, nil
}

func (svc *service) ListGoods(workspace string) ([]Good, DescribedError) {
	result, err := svc.db.GetAllItems(workspace)
	if err != nil {
		log.Error().Err(&ErrDBListFailed{err}).Msg("")
		return []Good{}, &ErrDBListFailed{err}
	}
	return result, nil
}

func (svc *service) DeleteGood(good Good) DescribedError {
	err := svc.db.DeleteItem(good)
	if err != nil {
		log.Error().Err(&ErrDBDeleteFailed{err}).Msg("")
		return &ErrDBDeleteFailed{err}
	}
	return nil
}

func (svc *service) BuildShoppingList(workspace string) ([]Good, DescribedError) {
	shoppingList := []Good{}

	result, err := svc.db.GetAllItems(workspace)
	if err != nil {
		log.Error().Err(&ErrShopDBListFailed{err}).Msg("")
		return []Good{}, &ErrShopDBListFailed{err}
	}

	for _, good := range result {
		if good.Quantity == "Empty" || good.Quantity == "Low" {
			shoppingList = append(shoppingList, good)
			continue
		}

		if good.OpenExpire != "" {
			good.Expire = good.OpenExpire
		}
		expired, err := svc.isExpired(good.Expire)
		if err != nil {
			log.Error().Err(&ErrShopParseDateFailed{err}).Msg("")
			return []Good{}, &ErrShopParseDateFailed{err}
		}
		if expired {
			shoppingList = append(shoppingList, good)
		}
	}
	return shoppingList, nil
}

func (svc *service) isExpired(date string) (bool, error) {
	dateFormat := "02/01/2006"
	currentDate := time.Now()

	parsedDate, err := time.Parse(dateFormat, date)
	if err != nil {
		return false, err
	}

	return parsedDate.Sub(currentDate) < 0, nil
}
