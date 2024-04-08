package core

//-------------------------------------------------------------------------------
// Principal
//-------------------------------------------------------------------------------

type User struct {
	Username       string
	GivenPassword  string
	ActualPassword string
}

//-------------------------------------------------------------------------------
// Swagger only
//-------------------------------------------------------------------------------

type UserLogin struct {
	Username string `json:"username" binding:"required" example:"john.foo"`
	Password string `json:"password" binding:"required" example:"Qwerty"`
}
