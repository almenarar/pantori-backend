package core

import "fmt"

var technicalError = `oops! something went wrong on our end, please try your request again in a few minutes - reach support if persists`

type DescribedError interface {
	Error() string
	PublicMessage() string
}

//-------------------------------------------------------------------------------

type ErrInvalidLoginInput struct {
	Err error
}

func (e *ErrInvalidLoginInput) Error() string {
	return fmt.Sprintf("Failure in authenticate - incorrect user or password: %s", e.Err.Error())
}

func (e *ErrInvalidLoginInput) PublicMessage() string {
	return "incorrect user or password"
}

//-------------------------------------------------------------------------------

type ErrGetUserFailed struct {
	Err error
}

func (e *ErrGetUserFailed) Error() string {
	return fmt.Sprintf("Failure in db when authenticating: %s", e.Err.Error())
}

func (e *ErrGetUserFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrGenTokenFailed struct {
	Err error
}

func (e *ErrGenTokenFailed) Error() string {
	return fmt.Sprintf("Failure in token gen when authenticating: %s", e.Err.Error())
}

func (e *ErrGenTokenFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrEncryptPwdFailed struct {
	Err error
}

func (e *ErrEncryptPwdFailed) Error() string {
	return fmt.Sprintf("Failure in crypt when creating user: %s", e.Err.Error())
}

func (e *ErrEncryptPwdFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBCreateUserFailed struct {
	Err error
}

func (e *ErrDBCreateUserFailed) Error() string {
	return fmt.Sprintf("Failure in db when creating user: %s", e.Err.Error())
}

func (e *ErrDBCreateUserFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBListUserFailed struct {
	Err error
}

func (e *ErrDBListUserFailed) Error() string {
	return fmt.Sprintf("Failure in db when listing users: %s", e.Err.Error())
}

func (e *ErrDBListUserFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDBDeleteUserFailed struct {
	Err error
}

func (e *ErrDBDeleteUserFailed) Error() string {
	return fmt.Sprintf("Failure in db when deleting user: %s", e.Err.Error())
}

func (e *ErrDBDeleteUserFailed) PublicMessage() string {
	return technicalError
}

//-------------------------------------------------------------------------------
