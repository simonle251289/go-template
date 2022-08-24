package usermodels

import (
	"encoding/json"
	"fmt"
	"template/databases"
)

type UserEntity struct {
	databases.BaseModel `json:",inline"`
	UserName            string `json:"userName" gorm:"column:username;not null; varchar(100)"`
	Password            string `json:"password" gorm:"column:password;not null;varchar(255)"`
	FirstName           string `json:"firstName" gorm:"column:first_name;varchar(100)"`
	LastName            string `json:"LastName" gorm:"column:last_name;varchar(100)"`
}

func (*UserEntity) TableName() string {
	return "users"
}

type UserResponse struct {
	databases.BaseModel `json:",inline"`
	UserName            string `json:"userName" gorm:"column:username"`
	FirstName           string `json:"firstName" gorm:"column:first_name"`
	LastName            string `json:"LastName" gorm:"column:last_name"`
}

func (*UserResponse) TableName() string {
	entity := UserEntity{}
	return entity.TableName()
}

func (user *UserEntity) ToResponse() UserResponse {
	userJson, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err)
	}
	var userRes UserResponse
	err = json.Unmarshal(userJson, &userRes)
	if err != nil {
		fmt.Println(err)
	}
	return userRes
}
