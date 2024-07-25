package core_test

import (
	"bytes"
	"pantori/internal/auth/core"
	"pantori/internal/auth/mocks"

	"errors"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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
	ExpectedLog        string
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
			ExpectedLog:        "",
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
			ExpectedLog:        "",
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
			ExpectedError:      &core.ErrGetUserFailed{},
			ExpectedLog:        "Failure in db when authenticating",
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
			ExpectedLog:        "",
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
			ExpectedError:      &core.ErrGenTokenFailed{},
			ExpectedLog:        "Failure in token gen when authenticating:",
			ExpectedToken:      "",
			ExpectedInvocation: "-Get-CheckPwd-GenToken",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			var buf bytes.Buffer
			globalLogger := zerolog.New(&buf).With().Timestamp().Logger()
			log.Logger = globalLogger

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

			logOutput := buf.String()

			assert.Equal(testCase.ExpectedToken, token)
			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if testCase.ExpectedError != nil {
				assert.Contains(logOutput, testCase.ExpectedLog)
				assert.NotEmpty(err.PublicMessage())
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
	ExpectedLog        string
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
			ExpectedLog:        "",
			ExpectedInvocation: "-GetTime-EncryptPwd-Create",
		},
		{
			Description:        "successful run without workspace",
			InputUser:          core.User{},
			WhenDbErr:          nil,
			WhenEncryptPwdErr:  nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-GetTime-GenerateID-EncryptPwd-Create",
		},
		{
			Description:        "encrypt error",
			InputUser:          core.User{},
			WhenDbErr:          nil,
			WhenEncryptPwdErr:  errors.New(""),
			ExpectedError:      &core.ErrEncryptPwdFailed{},
			ExpectedLog:        "Failure in crypt when creating user",
			ExpectedInvocation: "-GetTime-GenerateID-EncryptPwd",
		},
		{
			Description:        "db error",
			InputUser:          core.User{},
			WhenDbErr:          errors.New(""),
			WhenEncryptPwdErr:  nil,
			ExpectedError:      &core.ErrDBCreateUserFailed{},
			ExpectedLog:        "Failure in db when creating user",
			ExpectedInvocation: "-GetTime-GenerateID-EncryptPwd-Create",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			var buf bytes.Buffer
			globalLogger := zerolog.New(&buf).With().Timestamp().Logger()
			log.Logger = globalLogger

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
			logOutput := buf.String()

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if testCase.ExpectedError != nil {
				assert.Contains(logOutput, testCase.ExpectedLog)
				assert.NotEmpty(err.PublicMessage())
			}
		})
	}
}

type DeleteUserCase struct {
	Description        string
	InputUser          core.User
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedInvocation string
}

func TestDeleteUser(t *testing.T) {
	testCases := []DeleteUserCase{
		{
			Description:        "successful run",
			InputUser:          core.User{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-Delete",
		},
		{
			Description:        "db error",
			InputUser:          core.User{},
			WhenDbErr:          errors.New(""),
			ExpectedError:      &core.ErrDBDeleteUserFailed{},
			ExpectedLog:        "Failure in db when deleting user:",
			ExpectedInvocation: "-Delete",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			var buf bytes.Buffer
			globalLogger := zerolog.New(&buf).With().Timestamp().Logger()
			log.Logger = globalLogger

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
			logOutput := buf.String()

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if testCase.ExpectedError != nil {
				assert.Contains(logOutput, testCase.ExpectedLog)
				assert.NotEmpty(err.PublicMessage())
			}
		})
	}
}

type ListUsersCase struct {
	Description        string
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedInvocation string
}

func TestListUsers(t *testing.T) {
	testCases := []ListUsersCase{
		{
			Description:        "successful run",
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-List",
		},
		{
			Description:        "db error",
			WhenDbErr:          errors.New(""),
			ExpectedError:      &core.ErrDBListUserFailed{},
			ExpectedLog:        "Failure in db when listing users:",
			ExpectedInvocation: "-List",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			var buf bytes.Buffer
			globalLogger := zerolog.New(&buf).With().Timestamp().Logger()
			log.Logger = globalLogger

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
			logOutput := buf.String()

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if testCase.ExpectedError != nil {
				assert.Contains(logOutput, testCase.ExpectedLog)
				assert.NotEmpty(err.PublicMessage())
			}
		})
	}
}
