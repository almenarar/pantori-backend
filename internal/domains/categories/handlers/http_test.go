package handlers_test

import (
	"pantori/internal/domains/categories/core"
	"pantori/internal/domains/categories/handlers"
	"pantori/internal/domains/categories/mocks"

	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestHttpCase struct {
	Description             string
	Payload                 []byte
	ValidateCategoryPayload func(core.Category) error
	ValidateStrPayload      func(string) error
	ShouldBeInvoked         bool
	WhenError               error
	Expected                int
}

func TestCreateCategory(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(`{"Name":"foo","Color":"bar"}`),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				if g.Name != "foo" ||
					g.Color != "bar" ||
					g.Workspace != "test" {
					t.Fatalf("unexpected input: %s", g)
				}
				return nil
			},
		},
		{
			Description:        "dryrun",
			Payload:            []byte(`{"Name":"foo","Color":"bar"}`),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
		{
			Description:        "server error",
			Payload:            []byte(`{"Name":"foo","Color":"bar"}`),
			ShouldBeInvoked:    true,
			WhenError:          errors.New(""),
			Expected:           500,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
		{
			Description:        "invalid payload",
			Payload:            []byte(`{"testing":"testing"}`),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           400,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:      testCase.ValidateStrPayload,
				AnalyzeCategory: testCase.ValidateCategoryPayload,
				Err:             testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Set("workspace", "test")
			c.Request, err = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.CreateCategory(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestEditCategory(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(`{"ID":"zerp","Name":"foo","Color":"bar"}`),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				if g.Name != "foo" ||
					g.Color != "bar" ||
					g.ID != "zerp" ||
					g.Workspace != "test" {
					t.Fatalf("unexpected input: %s", g)
				}
				return nil
			},
		},
		{
			Description:        "dryrun",
			Payload:            []byte(`{"ID":"zerp","Name":"foo","Color":"bar"}`),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
		{
			Description:        "server error",
			Payload:            []byte(`{"ID":"zerp","Name":"foo","Color":"bar"}`),
			ShouldBeInvoked:    true,
			WhenError:          errors.New(""),
			Expected:           500,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
		{
			Description:        "invalid payload",
			Payload:            []byte(`{"testing":"testing"}`),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           400,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:      testCase.ValidateStrPayload,
				AnalyzeCategory: testCase.ValidateCategoryPayload,
				Err:             testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Set("workspace", "test")
			c.Request, err = http.NewRequest(http.MethodPatch, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.EditCategory(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestDeleteCategory(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(`{"ID":"zerp"}`),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				if g.ID != "zerp" ||
					g.Workspace != "test" {
					t.Fatalf("unexpected input: %s", g)
				}
				return nil
			},
		},
		{
			Description:        "dryrun",
			Payload:            []byte(`{"ID":"zerp"}`),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
		{
			Description:        "server error",
			Payload:            []byte(`{"ID":"zerp"}`),
			ShouldBeInvoked:    true,
			WhenError:          errors.New(""),
			Expected:           500,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
		{
			Description:        "invalid payload",
			Payload:            []byte(`{"testing":"testing"}`),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           400,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:      testCase.ValidateStrPayload,
				AnalyzeCategory: testCase.ValidateCategoryPayload,
				Err:             testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Set("workspace", "test")
			c.Request, err = http.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.DeleteCategory(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestCreateDefaultCategories(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(``),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				if g.Workspace != "test" {
					t.Fatalf("unexpected input: %s", g)
				}
				return nil
			},
		},
		{
			Description:        "dryrun",
			Payload:            []byte(``),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
		{
			Description:        "server error",
			Payload:            []byte(``),
			ShouldBeInvoked:    true,
			WhenError:          errors.New(""),
			Expected:           500,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:      testCase.ValidateStrPayload,
				AnalyzeCategory: testCase.ValidateCategoryPayload,
				Err:             testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Set("workspace", "test")
			c.Request, err = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.CreateDefaultCategories(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestListCategories(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(``),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				if g.Workspace != "test" {
					t.Fatalf("unexpected input: %s", g)
				}
				return nil
			},
		},
		{
			Description:        "dryrun",
			Payload:            []byte(``),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
		{
			Description:        "server error",
			Payload:            []byte(``),
			ShouldBeInvoked:    true,
			WhenError:          errors.New(""),
			Expected:           500,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateCategoryPayload: func(g core.Category) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:      testCase.ValidateStrPayload,
				AnalyzeCategory: testCase.ValidateCategoryPayload,
				Err:             testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Set("workspace", "test")
			c.Request, err = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.ListCategories(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}
