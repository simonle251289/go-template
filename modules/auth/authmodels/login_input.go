package authmodels

type LoginInput struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type RegisterInput struct {
	UserName  string `json:"userName"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
