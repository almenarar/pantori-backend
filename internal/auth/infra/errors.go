package infra

import (
	"fmt"
)

//-------------------------------------------------------------------------------

type ErrUserNotFound struct{}

func (e *ErrUserNotFound) Error() string {
	return "user not found in db"
}

//-------------------------------------------------------------------------------

type ErrDbOpFailed struct {
	Err error
}

func (e *ErrDbOpFailed) Error() string {
	return fmt.Sprintf("Database op failed: %s", e.Err.Error())
}

//-------------------------------------------------------------------------------

type ErrDataTransformFailed struct {
	Err error
}

func (e *ErrDataTransformFailed) Error() string {
	return fmt.Sprintf("convertion between db and app types failed: %s", e.Err.Error())
}

//-------------------------------------------------------------------------------

type ErrEncryptPwdFailed struct {
	Err error
}

func (e *ErrEncryptPwdFailed) Error() string {
	return fmt.Sprintf("Something wrong while encrypting password: %s", e.Err.Error())
}

//-------------------------------------------------------------------------------

type ErrCheckPwdFailed struct {
	Err error
}

func (e *ErrCheckPwdFailed) Error() string {
	return fmt.Sprintf("Something wrong while comparing hash and password: %s", e.Err.Error())
}

//-------------------------------------------------------------------------------

type ErrGenTokenFailed struct {
	Err error
}

func (e *ErrGenTokenFailed) Error() string {
	return fmt.Sprintf("Something wrong while generating access token: %s", e.Err.Error())
}

//-------------------------------------------------------------------------------
