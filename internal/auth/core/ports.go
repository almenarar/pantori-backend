package core

type ServicePort interface {
	Authenticate(User) (string, error)
}

type DatabasePort interface {
	GetUser(User) (User, error)
}

type CryptographyPort interface {
	CheckPassword(stored, given string) error
	GenerateToken(User) (string, error)
}
