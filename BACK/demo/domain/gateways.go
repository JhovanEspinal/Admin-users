package domain

import "demo/domain/model"

type DatabaseGateway interface {

	SaveUser(user *model.User)(*model.User,error)
	ValidateUser(registeredUser *model.RegisteredUser, user *model.User)(*model.User, error)
	RetrieveUsers()(*[]model.User,error)
	Setup()
	DeleteUser(id string)error
	UpdateUser(user *model.User)(*model.User,error)

}