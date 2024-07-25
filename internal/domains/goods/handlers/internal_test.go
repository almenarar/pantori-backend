package handlers_test

import (
	"pantori/internal/domains/goods/core"
	"pantori/internal/domains/goods/handlers"
	"pantori/internal/domains/goods/mocks"
	notifiers "pantori/internal/domains/notifiers/core"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListAllUsersCase struct {
	Description    string
	Invoked        bool
	WhenListErr    core.DescribedError
	ExpectedOutput []notifiers.Good
	ExpectedError  core.DescribedError
}

func TestListAllUsers(t *testing.T) {
	testCases := []ListAllUsersCase{
		{
			Description:    "successfull run",
			Invoked:        true,
			WhenListErr:    nil,
			ExpectedOutput: []notifiers.Good{{Name: "carrot"}, {Name: "grapes", Expire: "2000"}},
			ExpectedError:  nil,
		},
		{
			Description:    "fail at list goods",
			Invoked:        true,
			WhenListErr:    &core.ErrDBListFailed{},
			ExpectedOutput: []notifiers.Good(nil),
			ExpectedError:  &core.ErrDBListFailed{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{Err: testCase.WhenListErr, AnalyzeStr: func(s string) core.DescribedError { return nil }}
			internal := handlers.NewInternal(
				&svc,
			)

			out, err := internal.GetGoodsFromWorkspace("wkp")
			assert.IsType(testCase.ExpectedError, err)
			assert.Equal(testCase.Invoked, svc.Invoked)
			assert.Equal(out, testCase.ExpectedOutput)
		})
	}
}
