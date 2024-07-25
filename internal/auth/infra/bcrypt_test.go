package infra_test

import (
	"fmt"
	"pantori/internal/auth/core"
	"pantori/internal/auth/infra"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

type TestBcrypt struct {
	Description    string
	CreatePassword string
	GivenPassword  string
	ErrMsg         string
	ErrEncrypt     error
	ErrCheck       error
}

func TestPasswordFlow(t *testing.T) {
	testCases := []TestBcrypt{
		{
			Description:    "successfull run",
			CreatePassword: "secret",
			GivenPassword:  "secret",
			ErrMsg:         "",
			ErrEncrypt:     nil,
			ErrCheck:       nil,
		},
		{
			Description:    "wrong password",
			CreatePassword: "secret",
			GivenPassword:  "leaked",
			ErrMsg:         "Something wrong while comparing hash and password",
			ErrEncrypt:     nil,
			ErrCheck:       &infra.ErrCheckPwdFailed{},
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
			if err != nil {
				assert.Contains(err.Error(), testCase.ErrMsg)
			}
		})
	}
}

type TestJWT struct {
	Description string
	Input       core.User
}

func TestJWTFlow(t *testing.T) {
	testCases := []TestJWT{
		{
			Description: "successfull run",
			Input: core.User{
				Username:  "john",
				Workspace: "wkp1",
			},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			crypto := infra.NewCryptography("key")

			token, err := crypto.GenerateToken(testCase.Input)
			assert.IsType(nil, err)
			parsedToken, err := jwt.Parse(token, keyFunc)
			assert.IsType(nil, err)
			assert.True(parsedToken.Valid)
			claims, _ := parsedToken.Claims.(jwt.MapClaims)
			assert.Equal(testCase.Input.Username, claims["sub"].(string))
			assert.Equal(testCase.Input.Workspace, claims["workspace"].(string))
		})
	}
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte("key"), nil
}
