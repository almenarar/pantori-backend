package core

type Good struct {
	Name      string
	Workspace string
	Expire    string
}

type User struct {
	Name      string
	Email     string
	Workspace string
}

type Report struct {
	Username     string
	ExpiresToday []Good
	ExpiresSoon  []Good
	Expired      []Good
}
