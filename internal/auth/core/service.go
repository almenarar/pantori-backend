package core

type service struct {
	db     DatabasePort
	crypto CryptographyPort
}

func NewService(crypto CryptographyPort, db DatabasePort) *service {
	return &service{
		db:     db,
		crypto: crypto,
	}
}

func (svc *service) Authenticate(user User) (string, error) {
	user, err := svc.db.GetUser(user)
	if err != nil {
		return "", &ErrDbOpFailed{err}
	}

	// user not found
	if user.ActualPassword == "" {
		return "", &ErrInvalidLoginInput{}
	}

	err = svc.crypto.CheckPassword(user.ActualPassword, user.GivenPassword)
	if err != nil {
		return "", &ErrInvalidLoginInput{}
	}

	token, err := svc.crypto.GenerateToken(user)
	if err != nil {
		return "", &ErrCryptoOpFailed{err}
	}
	return token, nil
}
