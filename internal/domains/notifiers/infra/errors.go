package infra

import "fmt"

//-------------------------------------------------------------------------------

type ErrParseTemplate struct {
	Err error
}

func (r *ErrParseTemplate) Error() string {
	return fmt.Sprintf("Something wrong while parsing the html template: %s", r.Err.Error())
}

//-------------------------------------------------------------------------------

type ErrFillTemplate struct {
	Err error
}

func (r *ErrFillTemplate) Error() string {
	return fmt.Sprintf("Something wrong while filling the html template: %s", r.Err.Error())
}

//-------------------------------------------------------------------------------

type ErrSendEmail struct {
	Err error
}

func (r *ErrSendEmail) Error() string {
	return fmt.Sprintf("Something wrong while sending the email: %s", r.Err.Error())
}

//-------------------------------------------------------------------------------
