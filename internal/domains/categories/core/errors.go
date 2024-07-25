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
	return fmt.Sprintf("Failure in db when creating category: %s", e.Err.Error())
}

func (e *ErrDBCreateFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBListFailed struct {
	Err error
}

func (e *ErrDBListFailed) Error() string {
	return fmt.Sprintf("Failure in db when listing category: %s", e.Err.Error())
}

func (e *ErrDBListFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBEditFailed struct {
	Err error
}

func (e *ErrDBEditFailed) Error() string {
	return fmt.Sprintf("Failure in db when editing category: %s", e.Err.Error())
}

func (e *ErrDBEditFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBDeleteFailed struct {
	Err error
}

func (e *ErrDBDeleteFailed) Error() string {
	return fmt.Sprintf("Failure in db when deleting category: %s", e.Err.Error())
}

func (e *ErrDBDeleteFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------
