package mocks

import "pantori/internal/domains/notifiers/core"

type UsersMock struct {
	ErrList    error
	Invocation *string
}

func (um *UsersMock) ListAllUsersByWorkspace() (map[string][]core.User, error) {
	output := make(map[string][]core.User)
	*um.Invocation = *um.Invocation + "-ListUsers"
	if um.ErrList != nil {
		return output, um.ErrList
	}

	output["workspace1"] = []core.User{{
		Name: "user1",
	}}
	output["workspace2"] = []core.User{{
		Name: "user2",
	}}
	return output, nil
}
