package core

import "fmt"

var technicalError = `oops! something went wrong on our end, please try your request again in a few minutes - reach support if persists`

type DescribedError interface {
	Error() string
	PublicMessage() string
}

//-------------------------------------------------------------------------------

type ErrDBCreateFailed struct {
	Err error
}

func (e *ErrDBCreateFailed) Error() string {
	return fmt.Sprintf("Failure in db when creating good: %s", e.Err.Error())
}

func (e *ErrDBCreateFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBEditFailed struct {
	Err error
}

func (e *ErrDBEditFailed) Error() string {
	return fmt.Sprintf("Failure in db when editing good: %s", e.Err.Error())
}

func (e *ErrDBEditFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBListFailed struct {
	Err error
}

func (e *ErrDBListFailed) Error() string {
	return fmt.Sprintf("Failure in db when listing goods: %s", e.Err.Error())
}

func (e *ErrDBListFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBGetFailed struct {
	Err error
}

func (e *ErrDBGetFailed) Error() string {
	return fmt.Sprintf("Failure in db when getting goods: %s", e.Err.Error())
}

func (e *ErrDBGetFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBDeleteFailed struct {
	Err error
}

func (e *ErrDBDeleteFailed) Error() string {
	return fmt.Sprintf("Failure in db when deleting goods: %s", e.Err.Error())
}

func (e *ErrDBDeleteFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrShopDBListFailed struct {
	Err error
}

func (e *ErrShopDBListFailed) Error() string {
	return fmt.Sprintf("Failure in db when listing goods for shopping list: %s", e.Err.Error())
}

func (e *ErrShopDBListFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrShopParseDateFailed struct {
	Err error
}

func (e *ErrShopParseDateFailed) Error() string {
	return fmt.Sprintf("Failure parsing date when building shopping list: %s", e.Err.Error())
}

func (e *ErrShopParseDateFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------
