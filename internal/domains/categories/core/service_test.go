package core_test

import (
	"bytes"
	"errors"
	"pantori/internal/domains/categories/core"
	"pantori/internal/domains/categories/mocks"
	"strings"

	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

type AddCase struct {
	Description        string
	InputCategory      core.Category
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedInvocation string
}

func TestAdd(t *testing.T) {
	testCases := []AddCase{
		{
			Description:        "successfull add",
			InputCategory:      core.Category{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-GenerateID-GetTime-Add",
		},
		{
			Description:        "db error",
			InputCategory:      core.Category{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDBCreateFailed{},
			ExpectedLog:        "Failure in db when creating category:",
			ExpectedInvocation: "-GenerateID-GetTime-Add",
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
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.CreateCategory(testCase.InputCategory)
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

type AddDefaultCase struct {
	Description               string
	Input                     string
	WhenDbErr                 error
	ExpectedLastCategoryName  string
	ExpectedLastCategoryColor string
	ExpectedError             error
	ExpectedLog               string
	ExpectedInvocation        string
}

func TestDefaultAdd(t *testing.T) {
	testCases := []AddDefaultCase{
		{
			Description:               "successfull add",
			Input:                     "workspace",
			WhenDbErr:                 nil,
			ExpectedLastCategoryName:  "Outros",
			ExpectedLastCategoryColor: "FF92664A",
			ExpectedError:             nil,
			ExpectedLog:               "",
			ExpectedInvocation:        strings.Repeat("-GenerateID-GetTime-Add", 7),
		},
		{
			Description:               "db error",
			Input:                     "workspace",
			WhenDbErr:                 errors.New("some error"),
			ExpectedLastCategoryName:  "",
			ExpectedLastCategoryColor: "",
			ExpectedError:             &core.ErrDBCreateFailed{},
			ExpectedLog:               "Failure in db when creating category:",
			ExpectedInvocation:        "-GenerateID-GetTime-Add",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			var invocationTrail string

			var buf bytes.Buffer
			globalLogger := zerolog.New(&buf).With().Timestamp().Logger()
			log.Logger = globalLogger

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
			logOutput := buf.String()

			assert.Equal(testCase.ExpectedInvocation, invocationTrail)
			assert.IsType(testCase.ExpectedError, err)

			if err == nil {
				assert.Len(db.CreatedItems, 7)
				assert.Equal(testCase.ExpectedLastCategoryName, db.CreatedItems[6].Name)
				assert.Equal(testCase.ExpectedLastCategoryColor, db.CreatedItems[6].Color)
			} else {
				assert.Contains(logOutput, testCase.ExpectedLog)
				assert.NotEmpty(err.PublicMessage())
			}
		})
	}
}

type EditCase struct {
	Description        string
	InputCategory      core.Category
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedInvocation string
}

func TestEdit(t *testing.T) {
	testCases := []EditCase{
		{
			Description:        "successfull edit",
			InputCategory:      core.Category{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-GetTime-Edit",
		},
		{
			Description:        "db error",
			InputCategory:      core.Category{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDBEditFailed{},
			ExpectedLog:        "Failure in db when editing category:",
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
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.EditCategory(testCase.InputCategory)
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

type DeleteCase struct {
	Description        string
	InputCategory      core.Category
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedInvocation string
}

func TestDelete(t *testing.T) {
	testCases := []DeleteCase{
		{
			Description:        "successfull delete",
			InputCategory:      core.Category{},
			WhenDbErr:          nil,
			ExpectedError:      nil,
			ExpectedLog:        "",
			ExpectedInvocation: "-Delete",
		},
		{
			Description:        "db error",
			InputCategory:      core.Category{},
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDBDeleteFailed{},
			ExpectedLog:        "Failure in db when deleting category:",
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
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			err := svc.DeleteCategory(testCase.InputCategory)
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

type ListCase struct {
	Description        string
	WhenDbErr          error
	ExpectedError      error
	ExpectedLog        string
	ExpectedReturn     []core.Category
	ExpectedInvocation string
}

func TestList(t *testing.T) {
	testCases := []ListCase{
		{
			Description:   "successfull list",
			WhenDbErr:     nil,
			ExpectedError: nil,
			ExpectedLog:   "",
			ExpectedReturn: []core.Category{
				{ID: "foo"},
				{ID: "bar"},
			},
			ExpectedInvocation: "-List",
		},
		{
			Description:        "db error",
			WhenDbErr:          errors.New("some error"),
			ExpectedError:      &core.ErrDBListFailed{},
			ExpectedLog:        "Failure in db when listing category:",
			ExpectedReturn:     []core.Category{},
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
				&mocks.UtilsMocks{
					Invocation: &invocationTrail,
				},
			)

			out, err := svc.ListCategories("workspace")
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
