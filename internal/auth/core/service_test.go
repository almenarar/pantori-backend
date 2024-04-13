package core_test

import (
	"pantori/internal/auth/core"
	"pantori/internal/auth/mocks"

	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AuthCase struct {
	Description        string
	InputUser          core.User
	UserInDb           bool
	WhenDbErr          error
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
			WhenDbErr:          nil,
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
			WhenDbErr:          nil,
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
			WhenDbErr:          errors.New("some error"),
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
			WhenDbErr:          nil,
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
			WhenDbErr:          nil,
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
					Err:        testCase.WhenDbErr,
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
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

type CreateUserCase struct {
	Description        string
	InputUser          core.User
	WhenDbErr          error
	WhenEncryptPwdErr  error
	ExpectedError      error
	ExpectedInvocation string
}

func TestCreateUser(t *testing.T) {
	testCases := []CreateUserCase{
		{
			Description:        "successful run with workspace",
			InputUser:          core.User{Workspace: "test"},
			WhenDbErr:          nil,
			WhenEncryptPwdErr:  nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-GetTime-EncryptPwd-Create",
		},
		{
			Description:        "successful run without workspace",
			InputUser:          core.User{},
			WhenDbErr:          nil,
			WhenEncryptPwdErr:  nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-GetTime-GenerateID-EncryptPwd-Create",
		},
		{
			Description:        "encrypt error",
			InputUser:          core.User{},
			WhenDbErr:          nil,
			WhenEncryptPwdErr:  errors.New(""),
			ExpectedError:      &core.ErrCryptoOpFailed{},
			ExpectedInvocation: "-GetTime-GenerateID-EncryptPwd",
		},
		{
			Description:        "db error",
			InputUser:          core.User{},
			WhenDbErr:          errors.New(""),
			WhenEncryptPwdErr:  nil,
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedInvocation: "-GetTime-GenerateID-EncryptPwd-Create",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			svc := core.NewService(
				&mocks.CryptoMock{
					ErrEncryptPwd: testCase.WhenEncryptPwdErr,
					Invocation:    &invocationTrail,
				},
				&mocks.DatabaseMock{
					Err:        testCase.WhenDbErr,
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.CreateUser(testCase.InputUser)

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}

type DeleteUserCase struct {
	Description        string
	InputUser          core.User
	WhenDbErr          error
	ExpectedError      error
	ExpectedInvocation string
}

func TestDeleteUser(t *testing.T) {
	testCases := []DeleteUserCase{
		{
			Description:        "successful run",
			InputUser:          core.User{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-Delete",
		},
		{
			Description:        "db error",
			InputUser:          core.User{},
			WhenDbErr:          errors.New(""),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedInvocation: "-Delete",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			svc := core.NewService(
				&mocks.CryptoMock{},
				&mocks.DatabaseMock{
					Err:        testCase.WhenDbErr,
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.DeleteUser(testCase.InputUser)

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}

type ListUsersCase struct {
	Description        string
	WhenDbErr          error
	ExpectedError      error
	ExpectedInvocation string
}

func TestListUsers(t *testing.T) {
	testCases := []ListUsersCase{
		{
			Description:        "successful run",
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-List",
		},
		{
			Description:        "db error",
			WhenDbErr:          errors.New(""),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedInvocation: "-List",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			svc := core.NewService(
				&mocks.CryptoMock{},
				&mocks.DatabaseMock{
					Err:        testCase.WhenDbErr,
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			_, err := svc.ListUsers()

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}
