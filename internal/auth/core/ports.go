package core

type ServicePort interface {
	Authenticate(User) (string, error)
	CreateUser(User) error
	DeleteUser(User) error
	ListUsers() ([]User, error)
}

type DatabasePort interface {
	GetUser(string) (User, error)
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
