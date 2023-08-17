package model

import "main/helper/repository"

type User struct {
	repository.BaseModel `json:"-"`
	Account              string `gorm:"uniqueIndex:idx_user_code" json:"account"`  // 帳號
	Password             string `gorm:"uniqueIndex:idx_user_code" json:"password"` //密碼
}

func (User) TableName() string {
	return "user"
}

type UserRepository interface {
	Creat(creat *User) error
	GetAll() ([]User, error)
	Get(id string) (User, error)
	UpdateAccount(id string, newAccount string) (User, error)
	UpdatePassword(id string, newPassword string) (User, error)
	DeleteBy(id string) error
	Login(login *User) (User, error)
}

type UserUsecase interface {
	Create(creat *User) error
	GetAll() ([]User, error)
	Get(id string) (User, error)
	UpdateAccount(id string, newAccount string) (User, error)
	UpdatePassword(id string, newPassword string) (User, error)
	DeleteBy(id string) error
	Login(login *User) (User, error)
}
