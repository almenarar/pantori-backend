package core

import (
	"github.com/rs/zerolog/log"
)

var technicalError = `oops! something went wrong on our end, please try your request again in a few minutes`

//-------------------------------------------------------------------------------

type ErrDbOpFailed struct {
	Inner error
}

func (r *ErrDbOpFailed) Error() string {
	log.Error().Stack().Err(r.Inner).Msg("")
	return technicalError
}

//-------------------------------------------------------------------------------

type ErrDateParseError struct {
	Inner error
}

func (r *ErrDateParseError) Error() string {
	log.Error().Stack().Err(r.Inner).Msg("")
	return technicalError
}
