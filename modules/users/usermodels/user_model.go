package usermodels

import "template/databases"

type UserEntity struct {
	databases.BaseModel
	UserName  string `json:"userName" gorm:"not null; varchar(100)"`
	Password  string `json:"password" gorm:"not null;varchar(255)"`
	FirstName string `json:"firstName" gorm:"varchar(100)"`
	LastName  string `json:"LastName" gorm:"varchar(100)"`
}

func (UserEntity) TableName() string {
	return "users"
}
