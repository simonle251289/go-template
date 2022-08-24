package authmodels

import (
	"template/modules/users/usermodels"
)

type LoginInput struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (login *LoginInput) ValidateLoginInput() bool {
	return !(len(login.UserName) == 0 || len(login.Password) == 0)
}

type RegisterInput struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LoginResponse struct {
	User  usermodels.UserResponse `json:"user"`
	Token Token                   `json:"token"`
}
