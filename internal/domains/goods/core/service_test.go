package core_test

import (
	"bytes"
	"errors"
	"pantori/internal/domains/goods/core"
	"pantori/internal/domains/goods/mocks"

	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

type AddCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedInvocation string
}

func TestAdd(t *testing.T) {
	testCases := []AddCase{
		{
			Description:        "successfull add",
			InputGood:          core.Good{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-GetImage-GenerateID-GetTime-Add",
		},
		{
			Description:        "db error",
			InputGood:          core.Good{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDBCreateFailed{},
			ExpectedLog:        "Failure in db when creating good:",
			ExpectedInvocation: "-GetImage-GenerateID-GetTime-Add",
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
				&mocks.DatabaseMock{
					ErrAdd:     testCase.WhenDbErr,
					Invocation: &invocationTrail,
				},
				&mocks.ImageMock{
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.AddGood(testCase.InputGood)
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

type EditCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedInvocation string
}

func TestEdit(t *testing.T) {
	testCases := []EditCase{
		{
			Description:        "successfull edit",
			InputGood:          core.Good{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-GetTime-Edit",
		},
		{
			Description:        "db error",
			InputGood:          core.Good{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDBEditFailed{},
			ExpectedLog:        "Failure in db when editing good:",
			ExpectedInvocation: "-GetTime-Edit",
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
				&mocks.DatabaseMock{
					ErrEdit:    testCase.WhenDbErr,
					Invocation: &invocationTrail,
				},
				&mocks.ImageMock{
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.EditGood(testCase.InputGood)
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

type GetCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedReturn     core.Good
	ExpectedInvocation string
}

func TestGet(t *testing.T) {
	testCases := []GetCase{
		{
			Description:        "successfull list",
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedReturn:     core.Good{ID: "foo"},
			ExpectedInvocation: "-Get",
		},
		{
			Description:        "db error",
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDBGetFailed{},
			ExpectedLog:        "Failure in db when getting goods:",
			ExpectedReturn:     core.Good{},
			ExpectedInvocation: "-Get",
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
				&mocks.DatabaseMock{
					ErrGet:     testCase.WhenDbErr,
					Invocation: &invocationTrail,
				},
				&mocks.ImageMock{
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			out, err := svc.GetGood(testCase.InputGood)
			logOutput := buf.String()

			assert.Equal(testCase.ExpectedReturn, out)
			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if testCase.ExpectedError != nil {
				assert.Contains(logOutput, testCase.ExpectedLog)
				assert.NotEmpty(err.PublicMessage())
			}
		})
	}
}

type ListCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedReturn     []core.Good
	ExpectedInvocation string
}

func TestList(t *testing.T) {
	testCases := []ListCase{
		{
			Description:   "successfull list",
			WhenDbErr:     nil,
			ExpectedError: nil,
			ExpectedLog:   "",
			ExpectedReturn: []core.Good{
				{ID: "foo"},
				{ID: "bar"},
			},
			ExpectedInvocation: "-List",
		},
		{
			Description:        "db error",
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDBListFailed{},
			ExpectedLog:        "Failure in db when listing goods:",
			ExpectedReturn:     []core.Good{},
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
				&mocks.DatabaseMock{
					ErrList:    testCase.WhenDbErr,
					Invocation: &invocationTrail,
				},
				&mocks.ImageMock{
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			out, err := svc.ListGoods(testCase.InputGood.Workspace)
			logOutput := buf.String()

			assert.Equal(testCase.ExpectedReturn, out)
			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if testCase.ExpectedError != nil {
				assert.Contains(logOutput, testCase.ExpectedLog)
				assert.NotEmpty(err.PublicMessage())
			}
		})
	}
}

type DeleteCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedInvocation string
}

func TestDelete(t *testing.T) {
	testCases := []DeleteCase{
		{
			Description:        "successfull delete",
			InputGood:          core.Good{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-Delete",
		},
		{
			Description:        "db error",
			InputGood:          core.Good{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDBDeleteFailed{},
			ExpectedLog:        "Failure in db when deleting goods:",
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
				&mocks.DatabaseMock{
					ErrDelete:  testCase.WhenDbErr,
					Invocation: &invocationTrail,
				},
				&mocks.ImageMock{
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.DeleteGood(testCase.InputGood)
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

type ShoppingListCase struct {
	Description        string
	ExistingItems      []core.Good
	WhenDbErr          error
	WhenDateFormatErr  error
	ExpectedError      error
	ExpectedLog        string
	ExpectedInvocation string
	ExpectedOutput     []core.Good
}

func TestShoppingList(t *testing.T) {
	testCases := []ShoppingListCase{
		{
			Description: "empty list, no errors",
			ExistingItems: []core.Good{
				{
					Name:     "carrot",
					Expire:   "20/10/2080",
					Quantity: "Full",
				},
				{
					Name:     "grapes",
					Expire:   "20/10/2050",
					Quantity: "Regular",
				},
			},
			WhenDbErr:          nil,
			WhenDateFormatErr:  nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-List",
			ExpectedOutput:     []core.Good{},
		},
		{
			Description: "list with items, no errors",
			ExistingItems: []core.Good{
				{
					Name:     "carrot",
					Expire:   "20/10/2080",
					Quantity: "Empty",
				},
				{
					Name:     "beer",
					Expire:   "20/10/2050",
					Quantity: "Regular",
				},
				{
					Name:     "grapes",
					Expire:   "20/10/2010",
					Quantity: "Low",
				},
				{
					Name:       "beans",
					Expire:     "20/10/2060",
					OpenExpire: "10/01/2000",
					Quantity:   "Regular",
				},
			},
			WhenDbErr:          nil,
			WhenDateFormatErr:  nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-List",
			ExpectedOutput: []core.Good{
				{
					Name:     "carrot",
					Expire:   "20/10/2080",
					Quantity: "Empty",
				},
				{
					Name:     "grapes",
					Expire:   "20/10/2010",
					Quantity: "Low",
				},
				{
					Name:       "beans",
					Expire:     "10/01/2000",
					OpenExpire: "10/01/2000",
					Quantity:   "Regular",
				},
			},
		},
		{
			Description:        "db error",
			WhenDbErr:          errors.New("some error"),
			WhenDateFormatErr:  nil,
			ExpectedError:      &core.ErrShopDBListFailed{},
			ExpectedLog:        "Failure in db when listing goods for shopping list:",
			ExpectedInvocation: "-List",
			ExpectedOutput:     []core.Good{},
		},
		{
			Description: "parse error",
			ExistingItems: []core.Good{
				{
					Expire: "40/40/0",
				}},
			WhenDbErr:          nil,
			WhenDateFormatErr:  errors.New("some error"),
			ExpectedError:      &core.ErrShopParseDateFailed{},
			ExpectedLog:        "Failure parsing date when building shopping list:",
			ExpectedInvocation: "-List",
			ExpectedOutput:     []core.Good{},
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
				&mocks.DatabaseMock{
					ListItemsOutput: testCase.ExistingItems,
					ErrList:         testCase.WhenDbErr,
					Invocation:      &invocationTrail,
				},
				&mocks.ImageMock{
					Invocation: &invocationTrail,
				},
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			output, err := svc.BuildShoppingList("workspace")
			logOutput := buf.String()

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.Equal(testCase.ExpectedOutput, output)
			assert.IsType(testCase.ExpectedError, err)
			if testCase.ExpectedError != nil {
				assert.Contains(logOutput, testCase.ExpectedLog)
				assert.NotEmpty(err.PublicMessage())
			}
		})
	}
}
