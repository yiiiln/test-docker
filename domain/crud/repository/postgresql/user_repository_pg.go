package postgresql

import (
	"gorm.io/gorm"
	"main/model"
)

type userRepositoryPGImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryPGImpl(DB *gorm.DB) model.UserRepository {
	return &userRepositoryPGImpl{
		DB: DB,
	}
}

func (repo *userRepositoryPGImpl) Creat(creat *model.User) error {
	return repo.DB.Debug().Create(creat).Error
}

func (repo *userRepositoryPGImpl) GetAll() ([]model.User, error) {
	var users []model.User
	err := repo.DB.Debug().Find(&users).Error
	return users, err
}

func (repo *userRepositoryPGImpl) Get(id string) (model.User, error) {
	var user model.User
	err := repo.DB.Debug().First(&user, "id = ?", id).Error
	return user, err
}

func (repo *userRepositoryPGImpl) UpdateAccount(id string, newAccount string) (model.User, error) {
	var user model.User
	err := repo.DB.Debug().Model(&user).Where("id = ?", id).Update("account", newAccount).Error
	return user, err
}

func (repo *userRepositoryPGImpl) UpdatePassword(id string, newPassword string) (model.User, error) {
	var user model.User
	err := repo.DB.Debug().Model(&user).Where("id = ?", id).Update("password", newPassword).Error
	return user, err
}

func (repo *userRepositoryPGImpl) DeleteBy(id string) error {
	//TODO implement me
	var user model.User
	return repo.DB.Debug().Where("id = ?", id).Delete(&user).Error
}

func (repo *userRepositoryPGImpl) Login(login *model.User) (model.User, error) {
	var user model.User
	err := repo.DB.Debug().First(&user, "account = ? AND password = ?", login.Account, login.Password).Error
	return user, err
}

/*
func (repo *userRepositoryPGImpl) FindBy(account string) (model.User, error) {
	var user model.User
	err := repo.DB.
		Select("id", "aluminum_code").
		Where("account = ?", account).
		Preload("UserMetadata").
		First(&user).Error
	return user, err
}

*/
