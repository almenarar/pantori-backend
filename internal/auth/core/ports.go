package core

type ServicePort interface {
	Authenticate(User) (string, DescribedError)
	CreateUser(User) DescribedError
	DeleteUser(User) DescribedError
	ListUsers() ([]User, DescribedError)
}

type DatabasePort interface {
	GetUser(User) (User, error)
	ListUsers() ([]User, error)
	CreateUser(User) error
	DeleteUser(User) error
}

type CryptographyPort interface {
	CheckPassword(stored, given string) error
	EncryptPassword(password string) (string, error)
	GenerateToken(User) (string, error)
}

type UtilsPort interface {
	GenerateID() string
	GetCurrentTime() string
}
