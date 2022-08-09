package authrepos

import (
	"errors"
	"template/components/appcontext"
	"template/databases"
	"template/modules/users/usermodels"
)

func Login(ctx appcontext.AppContext, userName string, password string) (*usermodels.UserEntity, error) {
	if userName == "admin" && password == "123456" {
		return &usermodels.UserEntity{
			BaseModel: databases.BaseModel{ID: "uuid.UUID{}.Get()"},
			FirstName: "Simon",
			LastName:  "Le",
		}, nil
	}
	return nil, errors.New("Your username or password incorrect")
}
