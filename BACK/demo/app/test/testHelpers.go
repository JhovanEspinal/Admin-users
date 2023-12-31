package test

import (
	"demo/domain/model"
	"github.com/stretchr/testify/mock"
)

type DbGateway struct {
	mock.Mock
}

func NewDbGateway () DbGateway {
	return DbGateway{}
}

func (db DbGateway) Setup(){
}

func (db DbGateway) SaveUser(user *model.User)(*model.User,error){
args := db.Called(user)
if args.Get(0) == nil {
	return nil, args.Error(1)
}
return args.Get(0).(*model.User), args.Error(1)

}

func (db DbGateway) ValidateUser(registerUser *model.RegisteredUser,user *model.User)(*model.User,error){
	args := db.Called(registerUser,user)
	if args.Get(0) == nil || args.Get(1) == nil {
		return nil, args.Error(1)
	}
	return args.Get(1).(*model.User), args.Error(1)
}

func (db DbGateway) RetrieveUsers()(*[]model.User,error){

	agrs := db.Called()
	if agrs.Get(0) == nil {
		return nil, agrs.Error(1)
	}
	return agrs.Get(0).(*[]model.User), agrs.Error(1)
}
func (db DbGateway) DeleteUser(id string)error{
	args := db.Called(id)
	return args.Error(0)
}
func (db DbGateway) UpdateUser(user *model.User)(*model.User,error){

	args := db.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}
