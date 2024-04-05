package core_test

import (
	"errors"
	"pantori/internal/domains/categories/core"
	"pantori/internal/domains/categories/mocks"
	"strings"

	"testing"

	"github.com/stretchr/testify/assert"
)

type AddCase struct {
	Description        string
	InputCategory      core.Category
	WhenDbErr          error
	ExpectedError      error
	ExpectedInvocation string
}

func TestAdd(t *testing.T) {
	testCases := []AddCase{
		{
			Description:        "successfull add",
			InputCategory:      core.Category{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-GenerateID-GetTime-Add",
		},
		{
			Description:        "db error",
			InputCategory:      core.Category{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedInvocation: "-GenerateID-GetTime-Add",
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
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.CreateCategory(testCase.InputCategory)

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}

type AddDefaultCase struct {
	Description               string
	Input                     string
	WhenDbErr                 error
	ExpectedLastCategoryName  string
	ExpectedLastCategoryColor string
	ExpectedError             error
	ExpectedInvocation        string
}

func TestDefaultAdd(t *testing.T) {
	testCases := []AddDefaultCase{
		{
			Description:               "successfull add",
			Input:                     "workspace",
			WhenDbErr:                 nil,
			ExpectedLastCategoryName:  "Outros",
			ExpectedLastCategoryColor: "0xFF92664A",
			ExpectedError:             nil,
			ExpectedInvocation:        strings.Repeat("-GenerateID-GetTime-Add", 7),
		},
		{
			Description:               "db error",
			Input:                     "workspace",
			WhenDbErr:                 errors.New("some error"),
			ExpectedLastCategoryName:  "",
			ExpectedLastCategoryColor: "",
			ExpectedError:             &core.ErrDbOpFailed{},
			ExpectedInvocation:        "-GenerateID-GetTime-Add",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string
			db := &mocks.DatabaseMock{
				ErrAdd:     testCase.WhenDbErr,
				Invocation: &invocationTrail,
			}

			svc := core.NewService(
				db,
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.CreateDefaultCategories(testCase.Input)

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			} else {
				assert.Len(db.CreatedItems, 7)
				assert.Equal(testCase.ExpectedLastCategoryName, db.CreatedItems[6].Name)
				assert.Equal(testCase.ExpectedLastCategoryColor, db.CreatedItems[6].Color)
			}
		})
	}
}

type EditCase struct {
	Description        string
	InputCategory      core.Category
	WhenDbErr          error
	ExpectedError      error
	ExpectedInvocation string
}

func TestEdit(t *testing.T) {
	testCases := []EditCase{
		{
			Description:        "successfull edit",
			InputCategory:      core.Category{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-GetTime-Edit",
		},
		{
			Description:        "db error",
			InputCategory:      core.Category{},
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
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.EditCategory(testCase.InputCategory)

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
	InputCategory      core.Category
	WhenDbErr          error
	ExpectedError      error
	ExpectedInvocation string
}

func TestDelete(t *testing.T) {
	testCases := []DeleteCase{
		{
			Description:        "successfull delete",
			InputCategory:      core.Category{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedInvocation: "-Delete",
		},
		{
			Description:        "db error",
			InputCategory:      core.Category{},
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
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.DeleteCategory(testCase.InputCategory)

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
	ExpectedReturn     []core.Category
	ExpectedInvocation string
}

func TestList(t *testing.T) {
	testCases := []ListCase{
		{
			Description:   "successfull list",
			WhenDbErr:     nil,
			ExpectedError: nil,
			ExpectedReturn: []core.Category{
				{ID: "foo"},
				{ID: "bar"},
			},
			ExpectedInvocation: "-List",
		},
		{
			Description:        "db error",
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDbOpFailed{},
			ExpectedReturn:     []core.Category{},
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
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			out, err := svc.ListCategories("workspace")

			assert.Equal(testCase.ExpectedReturn, out)
			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)
			if err != nil {
				assert.Equal(testCase.ExpectedError.Error(), err.Error())
			}
		})
	}
}
