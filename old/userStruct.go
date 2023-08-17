package old

import (
	"errors"
	"gorm.io/gorm"
)

//`json:"auth_account" binding:"required"`

type Customer struct {
	CustomerAccount  string `json:"cust_account"`  // 帳號
	CustomerPassword string `json:"cust_password"` // 密碼
}

//`gorm:"size:255; not null; unique" json:"UserAccount"`

type DbUser struct {
	gorm.Model
	UserAccount  string `gorm:"size:255; not null;" json:"UserAccount"`  // db_帳號
	UserPassword string `gorm:"size:255; not null;" json:"UserPassword"` // db_密碼
}

type UpdateAccount struct {
	NewAccount string `json:"NewAccount"` // 欲更新的帳號
}

type UpdatePassword struct {
	NewPassword string `json:"NewPassword"` // 欲更新的密碼
}

type UpdateUserStruct struct {
	ID          uint   `json:"UpdateID"`
	OldAccount  string `json:"OldAccount"`
	NewAccount  string `json:"NewAccount"`
	OldPassword string `json:"OldPassword"`
	NewPassword string `json:"NewPassword"`
}

//type PSQLRepository interface {
//	Create() (*DbUser, error)
//	ReadUserByAccount(user *DbUser) (*DbUser, error)
//	UpdateUser(data *UpdateUserStruct) (*UpdateUserStruct, error)
//	DeleteUser(user *DbUser) (*DbUser, error)
//}

func (user *DbUser) Create() (*DbUser, error) {
	indexErr := errors.New("index Error")
	if user.UserAccount != "" && user.UserPassword != "" {
		err := Database.Debug().Create(user).Error
		if err != nil {
			return &DbUser{}, err
		}
		return user, nil
	}
	return user, indexErr
}

func (user *DbUser) ReadUserByAccount() (*DbUser, error) {
	indexErr := errors.New("index Error")
	responseErr := errors.New("RowsAffected == 0")

	if user.UserAccount != "" && user.UserPassword != "" {
		search := Database.Debug().Where("user_account = ? AND user_password = ?", user.UserAccount, user.UserPassword).First(user) //first

		if search.Error != nil {
			return &DbUser{}, search.Error
		} else if search.RowsAffected == 0 {
			return &DbUser{}, responseErr
		}
		return user, nil
	}
	return &DbUser{}, indexErr
}

func (data *UpdateUserStruct) UpdateUser() (*UpdateUserStruct, error) {
	indexErr := errors.New("index Error")
	responseErr := errors.New("RowsAffected == 0")

	if data.NewAccount != "" {
		findData := Database.Debug().Model(&DbUser{}).Where("id = ?", data.ID).Update("user_account", data.NewAccount)
		if findData.Error != nil {
			return &UpdateUserStruct{}, findData.Error
		} else if findData.RowsAffected == 0 {
			return &UpdateUserStruct{}, responseErr
		}
		return data, nil
	}

	if data.NewPassword != "" {
		findData := Database.Debug().Model(&DbUser{}).Where("id = ?", data.ID).Update("user_password", data.NewPassword)
		if findData.Error != nil {
			return &UpdateUserStruct{}, findData.Error
		} else if findData.RowsAffected == 0 {
			return &UpdateUserStruct{}, responseErr
		}
		return data, nil
	}

	return &UpdateUserStruct{}, indexErr
}

func (user *DbUser) DeleteUser() (*DbUser, error) {
	if deleteErr := Database.Debug().Where("user_account = ?", user.UserAccount).Delete(&user).Error; deleteErr != nil {
		return &DbUser{}, deleteErr
	}
	return user, nil
}
