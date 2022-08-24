package userrepos

import (
	"template/databases"
	"template/modules/users/usermodels"
)

func (repo *UserRepos) FindById(userId string) (*usermodels.UserResponse, error) {
	var user usermodels.UserResponse
	db := repo.ctx.GetMainPostgresConnection()
	err := db.Where(&usermodels.UserResponse{
		BaseModel: databases.BaseModel{
			ID:        userId,
			IsDeleted: false,
			IsEnable:  true,
		},
	}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepos) FindByUsername(username string) (*usermodels.UserEntity, error) {
	var user usermodels.UserEntity
	db := repo.ctx.GetMainPostgresConnection()
	err := db.Where(&usermodels.UserEntity{
		BaseModel: databases.BaseModel{
			IsDeleted: false,
			IsEnable:  true,
		},
		UserName: username,
	}).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
