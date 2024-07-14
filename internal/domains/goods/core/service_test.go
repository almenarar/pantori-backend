package core_test

import (
	"errors"
	"pantori/internal/domains/goods/core"
	"pantori/internal/domains/goods/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
)

type AddCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedInvocation string
}

func TestAdd(t *testing.T) {
	testCases := []AddCase{
		{
			Description:        "successfull add",
			InputGood:          core.Good{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-GetImage-GenerateID-GetTime-Add",
		},
		{
			Description:        "db error",
			InputGood:          core.Good{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedInvocation: "-GetImage-GenerateID-GetTime-Add",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

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

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}

type EditCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedInvocation string
}

func TestEdit(t *testing.T) {
	testCases := []EditCase{
		{
			Description:        "successfull edit",
			InputGood:          core.Good{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-GetTime-Edit",
		},
		{
			Description:        "db error",
			InputGood:          core.Good{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedInvocation: "-GetTime-Edit",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

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

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}

type GetCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedReturn     core.Good
	ExpectedInvocation string
}

func TestGet(t *testing.T) {
	testCases := []GetCase{
		{
			Description:        "successfull list",
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedReturn:     core.Good{ID: "foo"},
			ExpectedInvocation: "-Get",
		},
		{
			Description:        "db error",
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedReturn:     core.Good{},
			ExpectedInvocation: "-Get",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

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

			assert.Equal(testCase.ExpectedReturn, out)
			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}

type ListCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedReturn     []core.Good
	ExpectedInvocation string
}

func TestList(t *testing.T) {
	testCases := []ListCase{
		{
			Description:   "successfull list",
			WhenDbErr:     nil,
			ExpectedError: nil,
			ExpectedReturn: []core.Good{
				{ID: "foo"},
				{ID: "bar"},
			},
			ExpectedInvocation: "-List",
		},
		{
			Description:        "db error",
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedReturn:     []core.Good{},
			ExpectedInvocation: "-List",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

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

			assert.Equal(testCase.ExpectedReturn, out)
			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}

type DeleteCase struct {
	Description        string
	InputGood          core.Good
	WhenDbErr          error
	ExpectedError      error
	ExpectedInvocation string
}

func TestDelete(t *testing.T) {
	testCases := []DeleteCase{
		{
			Description:        "successfull delete",
			InputGood:          core.Good{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-Delete",
		},
		{
			Description:        "db error",
			InputGood:          core.Good{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedInvocation: "-Delete",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

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

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
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
					Quantity: "Regular",
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
					Quantity: "Regular",
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
			ExpectedError:      &core.ErrDbOpFailed{},
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
			ExpectedError:      &core.ErrDateParseError{},
			ExpectedInvocation: "-List",
			ExpectedOutput:     []core.Good{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

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

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.Equal(testCase.ExpectedOutput, output)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}
