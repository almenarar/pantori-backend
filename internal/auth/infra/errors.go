package infra

import (
	"errors"

	"github.com/rs/zerolog/log"
)

//-------------------------------------------------------------------------------

type ErrUserNotFound struct{}

func (e *ErrUserNotFound) Error() string {
	log.Error().Stack().Err(errors.New("user not found in db")).Msg("")
	return "user not found in db"
}

//-------------------------------------------------------------------------------

type ErrDbOpFailed struct {
	Inner error
}

func (e *ErrDbOpFailed) Error() string {
	log.Error().Stack().Err(e.Inner).Msg("")
	return e.Inner.Error()
}

//-------------------------------------------------------------------------------

type ErrDataTransformFailed struct {
	Inner error
}

func (e *ErrDataTransformFailed) Error() string {
	log.Error().Stack().Err(e.Inner).Msg("")
	return e.Inner.Error()
}
