package dtos

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RegisterDTO struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type TokenAuth struct {
	Token string `json:"token"`
}
