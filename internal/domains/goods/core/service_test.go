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
			ExpectedInvocation: "-GetImage-Add",
		},
		{
			Description:        "db error",
			InputGood:          core.Good{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedInvocation: "-GetImage-Add",
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
			)

			out, err := svc.ListGoods()

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
