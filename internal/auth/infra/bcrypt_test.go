package infra_test

import (
	"errors"
	"pantori/internal/auth/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestBcrypt struct {
	Description    string
	CreatePassword string
	GivenPassword  string
	ErrEncrypt     error
	ErrCheck       error
}

func TestPasswordFlow(t *testing.T) {
	testCases := []TestBcrypt{
		{
			Description:    "successfull run",
			CreatePassword: "secret",
			GivenPassword:  "secret",
			ErrEncrypt:     nil,
			ErrCheck:       nil,
		},
		{
			Description:    "wrong password",
			CreatePassword: "secret",
			GivenPassword:  "leaked",
			ErrEncrypt:     nil,
			ErrCheck:       errors.New(""),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			crypto := infra.NewCryptography("key")

			stored, err := crypto.EncryptPassword(testCase.CreatePassword)
			assert.IsType(testCase.ErrEncrypt, err)
			err = crypto.CheckPassword(stored, testCase.GivenPassword)
			assert.IsType(testCase.ErrCheck, err)
		})
	}
}
