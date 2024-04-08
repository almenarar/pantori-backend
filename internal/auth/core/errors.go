package core

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

var technicalError = `oops! something went wrong on our end, please try your request again in a few minutes`

//-------------------------------------------------------------------------------

type ErrInvalidLoginInput struct{}

func (e *ErrInvalidLoginInput) Error() string {
	log.Error().Stack().Err(errors.New("incorrect user or password")).Msg("")
	return "incorrect user or password"
}

//-------------------------------------------------------------------------------

type ErrDbOpFailed struct {
	Inner error
}

func (e *ErrDbOpFailed) Error() string {
	log.Error().Stack().Err(e.Inner).Msg("")
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrCryptoOpFailed struct {
	Inner error
}

func (e *ErrCryptoOpFailed) Error() string {
	log.Error().Stack().Err(e.Inner).Msg("")
	return technicalError
}

//-------------------------------------------------------------------------------
