package handlers_test

import (
	"pantori/internal/domains/goods/core"
	"pantori/internal/domains/goods/handlers"
	"pantori/internal/domains/goods/mocks"

	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestHttpCase struct {
	Description         string
	Payload             []byte
	ValidateGoodPayload func(core.Good) error
	ValidateStrPayload  func(string) error
	ShouldBeInvoked     bool
	WhenError           error
	Expected            int
}

func TestCreateGood(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(`{"Name":"foo","Categories":["bar","zep"],"Workspace":"main","Expire":"20/12/2032","BuyDate":"24/12/2032"}`),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
				if g.Name != "foo" ||
					g.Categories[0] != "bar" ||
					g.Workspace != "main" ||
					g.Expire != "20/12/2032" ||
					g.BuyDate != "24/12/2032" {
					t.Fatalf("unexpected input: %s", g)
				}
				return nil
			},
		},
		{
			Description:        "dryrun",
			Payload:            []byte(`{"Name":"foo","Categories":["bar","zep"],"Workspace":"main","Expire":"20/12/2032","BuyDate":"24/12/2032"}`),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
				return nil
			},
		},
		{
			Description:        "server error",
			Payload:            []byte(`{"Name":"foo","Categories":["bar","zep"],"Workspace":"main","Expire":"20/12/2032","BuyDate":"24/12/2032"}`),
			ShouldBeInvoked:    true,
			WhenError:          errors.New(""),
			Expected:           500,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
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
			ValidateGoodPayload: func(g core.Good) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:  testCase.ValidateStrPayload,
				AnalyzeGood: testCase.ValidateGoodPayload,
				Err:         testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Set("workspace", "main")
			c.Request, err = http.NewRequest(http.MethodPost, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.CreateGood(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestEditGood(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(`{"ID":"01","Name":"foo","Categories":["bar","zep"],"ImageURL":"http://pov.com","Workspace":"main","Expire":"20/12/2032","BuyDate":"24/12/2032","CreatedAt":"30/01/1996"}`),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
				if g.ID != "01" ||
					g.Name != "foo" ||
					g.Categories[0] != "bar" ||
					g.ImageURL != "http://pov.com" ||
					g.Workspace != "main" ||
					g.Expire != "20/12/2032" ||
					g.BuyDate != "24/12/2032" ||
					g.CreatedAt != "30/01/1996" {
					t.Fatalf("unexpected input: %s", g)
				}
				return nil
			},
		},
		{
			Description:        "dryrun",
			Payload:            []byte(`{"ID":"01","Name":"foo","Categories":["bar","zep"],"ImageURL":"http://pov.com","Workspace":"main","Expire":"20/12/2032","BuyDate":"24/12/2032","CreatedAt":"30/01/1996"}`),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
				return nil
			},
		},
		{
			Description:        "server error",
			Payload:            []byte(`{"ID":"01","Name":"foo","Categories":["bar","zep"],"ImageURL":"http://pov.com","Workspace":"main","Expire":"20/12/2032","BuyDate":"24/12/2032","CreatedAt":"30/01/1996"}`),
			ShouldBeInvoked:    true,
			WhenError:          errors.New(""),
			Expected:           500,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
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
			ValidateGoodPayload: func(g core.Good) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:  testCase.ValidateStrPayload,
				AnalyzeGood: testCase.ValidateGoodPayload,
				Err:         testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Set("workspace", "main")
			c.Request, err = http.NewRequest(http.MethodPatch, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.EditGood(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestDeleteGood(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(`{"ID":"01"}`),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
				if g.ID != "01" ||
					g.Workspace != "main" {
					t.Fatalf("unexpected input: %s", g)
				}
				return nil
			},
		},
		{
			Description:        "dryrun",
			Payload:            []byte(`{"ID":"01"}`),
			ShouldBeInvoked:    false,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
				return nil
			},
		},
		{
			Description:        "server error",
			Payload:            []byte(`{"ID":"01"}`),
			ShouldBeInvoked:    true,
			WhenError:          errors.New(""),
			Expected:           500,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
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
			ValidateGoodPayload: func(g core.Good) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:  testCase.ValidateStrPayload,
				AnalyzeGood: testCase.ValidateGoodPayload,
				Err:         testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Set("workspace", "main")
			c.Request, err = http.NewRequest(http.MethodDelete, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.DeleteGood(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestGetGood(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(``),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
				if g.ID != "foo" {
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
			ValidateGoodPayload: func(g core.Good) error {
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
			ValidateGoodPayload: func(g core.Good) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:  testCase.ValidateStrPayload,
				AnalyzeGood: testCase.ValidateGoodPayload,
				Err:         testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Request, err = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer(testCase.Payload))
			c.AddParam("id", "foo")
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.GetGood(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}

func TestListGood(t *testing.T) {
	testCases := []TestHttpCase{
		{
			Description:        "successfull http request",
			Payload:            []byte(``),
			ShouldBeInvoked:    true,
			WhenError:          nil,
			Expected:           200,
			ValidateStrPayload: func(s string) error { return nil },
			ValidateGoodPayload: func(g core.Good) error {
				if g.ID != "01" {
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
			ValidateGoodPayload: func(g core.Good) error {
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
			ValidateGoodPayload: func(g core.Good) error {
				return nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{
				AnalyzeStr:  testCase.ValidateStrPayload,
				AnalyzeGood: testCase.ValidateGoodPayload,
				Err:         testCase.WhenError,
			}
			h := handlers.NewNetwork(&svc)

			var err error
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			if testCase.Description == "dryrun" {
				c.Set("username", "dryrun")
			}

			c.Request, err = http.NewRequest(http.MethodGet, "/", bytes.NewBuffer(testCase.Payload))
			if err != nil {
				t.Fatalf(err.Error())
			}

			h.ListGoods(c)
			assert.Equal(testCase.ShouldBeInvoked, svc.Invoked)
			assert.Equal(testCase.Expected, w.Code)
		})
	}
}
