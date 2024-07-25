package core

import "fmt"

//-------------------------------------------------------------------------------

type ErrTimeParse struct {
	Date string
}

func (r *ErrTimeParse) Error() string {
	return fmt.Sprintf("Given date have incorrect format: %s", r.Date)
}

//-------------------------------------------------------------------------------

type ErrListUsers struct {
	Err error
}

func (r *ErrListUsers) Error() string {
	return fmt.Sprintf("Something wrong while listing users: %s", r.Err.Error())
}

//-------------------------------------------------------------------------------

type ErrGetGoods struct {
	Err error
}

func (r *ErrGetGoods) Error() string {
	return fmt.Sprintf("Something wrong while getting goods: %s", r.Err.Error())
}

//-------------------------------------------------------------------------------

type ErrSendEmail struct {
	Err error
}

func (r *ErrSendEmail) Error() string {
	return fmt.Sprintf("Something wrong while sending email: %s", r.Err.Error())
}

//-------------------------------------------------------------------------------
