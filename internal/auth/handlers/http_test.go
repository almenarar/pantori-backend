package handlers_test

import (
	"errors"
	"pantori/internal/auth/core"
	"pantori/internal/auth/handlers"
	"pantori/internal/auth/mocks"

	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestHttpCase struct {
	Description         string
	Payload             []byte
	ValidatePayloadFunc func(core.User) core.DescribedError
	ShouldBeInvoked     bool
	WhenError           core.DescribedError
	Expected            int
}

func TestLogin(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:     "successfull http request",
			Payload:         []byte(`{"username":"foo","password":"bar"}`),
			ShouldBeInvoked: true,
			WhenError:       nil,
			Expected:        200,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				if u.Username != "foo" ||
					u.GivenPassword != "bar" {
					t.Fatalf("unexpected input: %s", u)
				}
				return nil
			},
		},
		{
			Description:     "service failed",
			Payload:         []byte(`{"username":"foo","password":"bar"}`),
			ShouldBeInvoked: true,
			WhenError:       &core.ErrGenTokenFailed{Err: errors.New("something")},
			Expected:        500,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				return nil
			},
		},
		{
			Description:     "invalid payload",
			Payload:         []byte(`{"testing":"testing"}`),
			ShouldBeInvoked: false,
			WhenError:       nil,
			Expected:        400,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				return nil
			},
		},
		{
			Description:     "wrong password",
			Payload:         []byte(`{"username":"foo","password":"bar"}`),
			ShouldBeInvoked: true,
			WhenError:       &core.ErrInvalidLoginInput{Err: errors.New("something")},
			Expected:        400,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				CustomFunc: testCase.ValidatePayloadFunc,
				Err:        testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, err = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.Login(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestCreateUser(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:     "successfull http request",
			Payload:         []byte(`{"username":"foo","password":"bar","workspace":"principal","email":"john.doe@mail.com"}`),
			ShouldBeInvoked: true,
			WhenError:       nil,
			Expected:        200,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				if u.Username != "foo" ||
					u.GivenPassword != "bar" ||
					u.Workspace != "principal" ||
					u.Email != "john.doe@mail.com" {
					t.Fatalf("unexpected input: %s", u)
				}
				return nil
			},
		},
		{
			Description:     "service failed",
			Payload:         []byte(`{"username":"foo","password":"bar","workspace":"principal","email":"john.doe@mail.com"}`),
			ShouldBeInvoked: true,
			WhenError:       &core.ErrDBCreateUserFailed{Err: errors.New("something")},
			Expected:        500,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				return nil
			},
		},
		{
			Description:     "invalid payload",
			Payload:         []byte(`{"testing":"testing"}`),
			ShouldBeInvoked: false,
			WhenError:       nil,
			Expected:        400,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				CustomFunc: testCase.ValidatePayloadFunc,
				Err:        testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, err = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.CreateUser(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:     "successfull http request",
			Payload:         []byte(`{"username":"foo"}`),
			ShouldBeInvoked: true,
			WhenError:       nil,
			Expected:        200,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				if u.Username != "foo" {
					t.Fatalf("unexpected input: %s", u)
				}
				return nil
			},
		},
		{
			Description:     "service failed",
			Payload:         []byte(`{"username":"foo"}`),
			ShouldBeInvoked: true,
			WhenError:       &core.ErrDBDeleteUserFailed{Err: errors.New("something")},
			Expected:        500,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				return nil
			},
		},
		{
			Description:     "invalid payload",
			Payload:         []byte(`{"testing":"testing"}`),
			ShouldBeInvoked: false,
			WhenError:       nil,
			Expected:        400,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				CustomFunc: testCase.ValidatePayloadFunc,
				Err:        testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, err = http.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.DeleteUser(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestListUsers(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:     "successfull http request",
			Payload:         []byte(``),
			ShouldBeInvoked: true,
			WhenError:       nil,
			Expected:        200,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				if u.Username != "foo" {
					t.Fatalf("unexpected input: %s", u)
				}
				return nil
			},
		},
		{
			Description:     "service failed",
			Payload:         []byte(``),
			ShouldBeInvoked: true,
			WhenError:       &core.ErrDBListUserFailed{Err: errors.New("something")},
			Expected:        500,
			ValidatePayloadFunc: func(u core.User) core.DescribedError {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				CustomFunc: testCase.ValidatePayloadFunc,
				Err:        testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, err = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.ListUsers(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}
