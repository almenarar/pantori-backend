package handlers_test

import (
	"pantori/internal/auth/core"
	"pantori/internal/auth/handlers"
	"pantori/internal/auth/mocks"

	notifiers "pantori/internal/domains/notifiers/core"

	"testing"

	"github.com/stretchr/testify/assert"
)

type ListAllUsersCase struct {
	Description    string
	Invoked        bool
	WhenListErr    core.DescribedError
	ExpectedOutput map[string][]notifiers.User
	ExpectedError  core.DescribedError
}

func TestListAllUsers(t *testing.T) {
	testCases := []ListAllUsersCase{
		{
			Description:    "successfull run",
			Invoked:        true,
			WhenListErr:    nil,
			ExpectedOutput: map[string][]notifiers.User{"wkp1": {{Name: "john", Workspace: "wkp1"}}},
			ExpectedError:  nil,
		},
		{
			Description:    "fail at list users",
			Invoked:        true,
			WhenListErr:    &core.ErrDBListUserFailed{},
			ExpectedOutput: map[string][]notifiers.User{},
			ExpectedError:  &core.ErrDBListUserFailed{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			assert := assert.New(t)
			svc := mocks.Service{Err: testCase.WhenListErr}
			internal := handlers.NewInternal(
				&svc,
			)

			out, err := internal.ListAllUsersByWorkspace()
			assert.IsType(testCase.ExpectedError, err)
			assert.Equal(testCase.Invoked, svc.Invoked)
			assert.Equal(out, testCase.ExpectedOutput)
		})
	}
}
