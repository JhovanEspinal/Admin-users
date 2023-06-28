package useCases

import (
	"demo/domain"
	"demo/domain/model"
	"errors"
)

type UserUseCase struct {

	datastore domain.DatabaseGateway

}

func NewUserUseCase (datastore domain.DatabaseGateway)UserUseCase{

	return UserUseCase{
		datastore: datastore,
	}
}

func (useCase UserUseCase) SaveUser(user *model.User)(*model.User,error){

	var (
		registeredUser model.RegisteredUser
		User model.User
	)
	registeredUser.Email,registeredUser.Password = user.Email, user.Password

	useCase.datastore.ValidateUser(&registeredUser,&User)
	if User.Email == user.Email{
		return nil, errors.New("REGISTERED_EMAIL")
	}

	return useCase.datastore.SaveUser(user)
}

func (useCase UserUseCase) ValidateUser(registeredUser *model.RegisteredUser,user *model.User)(*model.User,error){

	return useCase.datastore.ValidateUser(registeredUser,user)
}

func (useCase UserUseCase) RetrieveUser()(*[]model.User,error){

	return useCase.datastore.RetrieveUsers()
}

func (useCase UserUseCase) DeleteUser(id string) error{

	return useCase.datastore.DeleteUser(id)
}

func (useCase UserUseCase) UpdateUser(user *model.User)error{

_, err :=  useCase.datastore.UpdateUser(user)
	return err
}
