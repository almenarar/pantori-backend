package core

//-------------------------------------------------------------------------------
// Principal
//-------------------------------------------------------------------------------

type User struct {
	Username       string
	GivenPassword  string
	ActualPassword string
	Workspace      string
	Email          string
	LastSeen       string
	CreatedAt      string
}

//-------------------------------------------------------------------------------
// DB only
//-------------------------------------------------------------------------------

type UserDB struct {
	Username  string
	Password  string
	Workspace string
	Email     string
	LastSeen  string
	CreatedAt string
}

//-------------------------------------------------------------------------------
// Swagger only
//-------------------------------------------------------------------------------

type UserLogin struct {
	Username string `json:"username" binding:"required" example:"john.foo"`
	Password string `json:"password" binding:"required" example:"Qwerty"`
}

type CreateUser struct {
	Username  string `json:"username" binding:"required" example:"john.foo"`
	Password  string `json:"password" binding:"required" example:"qwerty"`
	Workspace string `json:"workspace" example:"principal"`
	Email     string `json:"email" binding:"required" example:"john.foo@mail.com"`
}

type DeleteUser struct {
	Username string `json:"username" binding:"required" example:"john.foo"`
}
