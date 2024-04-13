package core

type service struct {
	db     DatabasePort
	crypto CryptographyPort
	utils  UtilsPort
}

func NewService(crypto CryptographyPort, db DatabasePort, utils UtilsPort) *service {
	return &service{
		db:     db,
		crypto: crypto,
		utils:  utils,
	}
}

func (svc *service) Authenticate(input User) (string, error) {
	user, err := svc.db.GetUser(input.Username)
	if err != nil {
		return "", &ErrDbOpFailed{err}
	}

	if user.ActualPassword == "" {
		return "", &ErrInvalidLoginInput{}
	}

	err = svc.crypto.CheckPassword(user.ActualPassword, input.GivenPassword)
	if err != nil {
		return "", &ErrInvalidLoginInput{}
	}

	token, err := svc.crypto.GenerateToken(user)
	if err != nil {
		return "", &ErrCryptoOpFailed{err}
	}
	return token, nil
}

func (svc *service) CreateUser(user User) error {
	user.CreatedAt = svc.utils.GetCurrentTime()

	if user.Workspace == "" {
		user.Workspace = svc.utils.GenerateID()
	}

	var err error
	user.GivenPassword, err = svc.crypto.EncryptPassword(user.GivenPassword)
	if err != nil {
		return &ErrCryptoOpFailed{err}
	}

	err = svc.db.CreateUser(user)
	if err != nil {
		return &ErrDbOpFailed{err}
	}

	return nil
}

func (svc *service) DeleteUser(user User) error {
	err := svc.db.DeleteUser(user)
	if err != nil {
		return &ErrDbOpFailed{err}
	}
	return nil
}

func (svc *service) ListUsers() ([]User, error) {
	out, err := svc.db.ListUsers()
	if err != nil {
		return []User{}, &ErrDbOpFailed{err}
	}
	return out, nil
}
