package mocks

type UtilsMocks struct {
	Invocation *string
}

func (um *UtilsMocks) GenerateID() string {
	*um.Invocation = *um.Invocation + "-GenerateID"
	return "ID"
}

func (um *UtilsMocks) GetCurrentTime() string {
	*um.Invocation = *um.Invocation + "-GetTime"
	return "time"
}
