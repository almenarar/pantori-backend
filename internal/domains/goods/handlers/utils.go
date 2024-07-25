package handlers

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func formatValidationError(err error) []string {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		var errs []string
		for _, fe := range ve {
			errs = append(errs, fe.Field())
		}
		return errs
	}
	return nil
}
