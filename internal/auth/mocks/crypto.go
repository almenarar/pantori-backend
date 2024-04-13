package mocks

import core "pantori/internal/auth/core"

type CryptoMock struct {
	ErrCheckPwd   error
	ErrGenToken   error
	ErrEncryptPwd error
	Invocation    *string
}

func (cm *CryptoMock) CheckPassword(stored, given string) error {
	*cm.Invocation = *cm.Invocation + "-CheckPwd"
	if cm.ErrCheckPwd != nil {
		return cm.ErrCheckPwd
	}
	return nil
}

func (cm *CryptoMock) GenerateToken(core.User) (string, error) {
	*cm.Invocation = *cm.Invocation + "-GenToken"
	if cm.ErrGenToken != nil {
		return "", cm.ErrGenToken
	}
	return "token", nil
}

func (cm *CryptoMock) EncryptPassword(password string) (string, error) {
	*cm.Invocation = *cm.Invocation + "-EncryptPwd"
	if cm.ErrEncryptPwd != nil {
		return "", cm.ErrEncryptPwd
	}
	return "solarbean", nil
}
