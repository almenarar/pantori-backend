package authhdl

import (
	core "pantori/internal/auth/core"

	"net/http"
)

func defineHTTPStatus(err error) int {
	switch err.(type) {
	case *core.ErrInvalidLoginInput:
		// User error (HTTP 400 Bad Request)
		return http.StatusBadRequest
	}

	return http.StatusInternalServerError
}
