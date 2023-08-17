package usecase

import "main/model"

type userUsecase struct {
	userRepo model.UserRepository
}

func NewUserUsecase(repo model.UserRepository) model.UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

func (userUsecase userUsecase) Create(creat *model.User) error {
	//TODO implement me
	err := userUsecase.userRepo.Creat(creat)
	return err
}

func (userUsecase userUsecase) GetAll() ([]model.User, error) {
	//TODO implement me
	userSlice, err := userUsecase.userRepo.GetAll()
	return userSlice, err
}

func (userUsecase userUsecase) Get(id string) (model.User, error) {
	//TODO implement me
	user, err := userUsecase.userRepo.Get(id)
	return user, err
}

func (userUsecase userUsecase) UpdateAccount(id string, newAccount string) (model.User, error) {
	//TODO implement me
	user, err := userUsecase.userRepo.UpdateAccount(id, newAccount)
	return user, err
}

func (userUsecase userUsecase) UpdatePassword(id string, newPassword string) (model.User, error) {
	//TODO implement me
	user, err := userUsecase.userRepo.UpdatePassword(id, newPassword)
	return user, err
}

func (userUsecase userUsecase) DeleteBy(id string) error {
	//TODO implement me
	err := userUsecase.userRepo.DeleteBy(id)
	return err
}

func (userUsecase userUsecase) Login(login *model.User) (model.User, error) {
	//TODO implement me
	user, err := userUsecase.userRepo.Login(login)
	return user, err
}
