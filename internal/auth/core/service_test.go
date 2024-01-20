package authcore_test

import (
	core "pantori/internal/auth/core"
	mocks "pantori/internal/auth/mocks"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AuthCase struct {
	Description        string
	InputUser          core.User
	UserInDb           bool
	WhenGetUserErr     error
	WhenCheckPwdErr    error
	WhenGenTokenErr    error
	ExpectedToken      string
	ExpectedError      error
	ExpectedInvocation string
}

func TestAuthenticate(t *testing.T) {
	testCases := []AuthCase{
		{
			Description:        "successful run",
			InputUser:          core.User{},
			UserInDb:           true,
			WhenGetUserErr:     nil,
			WhenCheckPwdErr:    nil,
			WhenGenTokenErr:    nil,
			ExpectedError:      nil,
			ExpectedToken:      "token",
			ExpectedInvocation: "-Get-CheckPwd-GenToken",
		},
		{
			Description:        "user not found",
			InputUser:          core.User{},
			UserInDb:           false,
			WhenGetUserErr:     nil,
			WhenCheckPwdErr:    nil,
			WhenGenTokenErr:    nil,
			ExpectedError:      &core.ErrInvalidLoginInput{},
			ExpectedToken:      "",
			ExpectedInvocation: "-Get",
		},
		{
			Description:        "db error",
			InputUser:          core.User{},
			UserInDb:           false,
			WhenGetUserErr:     errors.New("some error"),
			WhenCheckPwdErr:    nil,
			WhenGenTokenErr:    nil,
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedToken:      "",
			ExpectedInvocation: "-Get",
		},
		{
			Description:        "check pwd error",
			InputUser:          core.User{},
			UserInDb:           true,
			WhenGetUserErr:     nil,
			WhenCheckPwdErr:    errors.New("some error"),
			WhenGenTokenErr:    nil,
			ExpectedError:      &core.ErrInvalidLoginInput{},
			ExpectedToken:      "",
			ExpectedInvocation: "-Get-CheckPwd",
		},
		{
			Description:        "gen token error",
			InputUser:          core.User{},
			UserInDb:           true,
			WhenGetUserErr:     nil,
			WhenCheckPwdErr:    nil,
			WhenGenTokenErr:    errors.New("some error"),
			ExpectedError:      &core.ErrCryptoOpFailed{},
			ExpectedToken:      "",
			ExpectedInvocation: "-Get-CheckPwd-GenToken",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			svc := core.NewService(
				&mocks.CryptoMock{
					ErrCheckPwd: testCase.WhenCheckPwdErr,
					ErrGenToken: testCase.WhenGenTokenErr,
					Invocation:  &invocationTrail,
				},
				&mocks.DatabaseMock{
					UserExists: testCase.UserInDb,
					ErrGet:     testCase.WhenGetUserErr,
					Invocation: &invocationTrail,
				},
			)

			token, err := svc.Authenticate(testCase.InputUser)

			assert.Equal(testCase.ExpectedToken, token)
			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}
