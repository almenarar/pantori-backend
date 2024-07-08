package mocks

import "pantori/internal/domains/notifiers/core"

type EmailMock struct {
	ErrSend    error
	Invocation *string
}

func (em *EmailMock) SendEmail(user core.User, report core.Report) error {
	*em.Invocation = *em.Invocation + "-SendEmail"
	if em.ErrSend != nil {
		return em.ErrSend
	}
	return nil
}
