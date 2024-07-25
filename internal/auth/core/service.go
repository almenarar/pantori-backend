package core

import "github.com/rs/zerolog/log"

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

func (svc *service) Authenticate(input User) (string, DescribedError) {
	user, err := svc.db.GetUser(input)
	if err != nil {
		log.Error().Err(&ErrGetUserFailed{err}).Msg("")
		return "", &ErrGetUserFailed{}
	}

	if user.ActualPassword == "" {
		return "", &ErrInvalidLoginInput{}
	}

	err = svc.crypto.CheckPassword(user.ActualPassword, input.GivenPassword)
	if err != nil {
		log.Error().Err(&ErrInvalidLoginInput{err}).Msg("")
		return "", &ErrInvalidLoginInput{}
	}

	token, err := svc.crypto.GenerateToken(user)
	if err != nil {
		log.Error().Err(&ErrGenTokenFailed{err}).Msg("")
		return "", &ErrGenTokenFailed{}
	}
	return token, nil
}

func (svc *service) CreateUser(user User) DescribedError {
	user.CreatedAt = svc.utils.GetCurrentTime()

	if user.Workspace == "" {
		user.Workspace = svc.utils.GenerateID()
	}

	var err error
	user.GivenPassword, err = svc.crypto.EncryptPassword(user.GivenPassword)
	if err != nil {
		log.Error().Err(&ErrEncryptPwdFailed{err}).Msg("")
		return &ErrEncryptPwdFailed{}
	}

	err = svc.db.CreateUser(user)
	if err != nil {
		log.Error().Err(&ErrDBCreateUserFailed{err}).Msg("")
		return &ErrDBCreateUserFailed{}
	}

	return nil
}

func (svc *service) DeleteUser(user User) DescribedError {
	err := svc.db.DeleteUser(user)
	if err != nil {
		log.Error().Err(&ErrDBDeleteUserFailed{err}).Msg("")
		return &ErrDBDeleteUserFailed{}
	}
	return nil
}

func (svc *service) ListUsers() ([]User, DescribedError) {
	out, err := svc.db.ListUsers()
	if err != nil {
		log.Error().Err(&ErrDBListUserFailed{err}).Msg("")
		return []User{}, &ErrDBListUserFailed{}
	}
	return out, nil
}
