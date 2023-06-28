package useCases

import (
	"demo/app/test"
	"demo/domain/model"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSaveUserIsOk(t *testing.T) {
  app:= test.CreateApp()

  user := &model.User{
  	Id: "xxx1234",
  	Name: "Tigre Maligno",
  	Job: "developer",
  	Description: "pruebas 2021",
  	Cc: "123456789",
  	Gender: "masculino",
  	Email: "tigressa@gmail.sofka.com.co",
  	Password: "ABC123@",
  }

  app.DataStore.On("SaveUser",mock.Anything).Return(user,nil)
  app.DataStore.On("ValidateUser",mock.Anything,mock.Anything).Return(user,nil)

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.SaveUser(user)

	 assert.Nil(t,err)
		app.DataStore.MethodCalled("SaveUser",mock.Anything)
		app.DataStore.MethodCalled("ValidateUser",mock.Anything)

}

func TestSaveUserIsWrongConecctionFailed(t *testing.T) {
	app:= test.CreateApp()

	user := &model.User{
		Id: "xxx1234",
		Name: "Tigre Maligno",
		Job: "developer",
		Description: "pruebas 2021",
		Cc: "123456789",
		Gender: "masculino",
		Email: "tigressa@gmail.sofka.com.co",
		Password: "ABC123@",
	}

	app.DataStore.On("SaveUser",mock.Anything).Return(nil,errors.New("CONNECTION_FAIL"))
	app.DataStore.On("ValidateUser",mock.Anything,mock.Anything).Return(nil,nil)

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.SaveUser(user)

	assert.NotNil(t,err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("SaveUser",mock.Anything)


}

func TestValidateUserIsOk(t *testing.T) {
	app:= test.CreateApp()

	user := &model.User{
		Id: "xxx1234",
		Name: "Tigre Maligno",
		Job: "developer",
		Description: "pruebas 2021",
		Cc: "123456789",
		Gender: "masculino",
		Email: "tigressa@gmail.sofka.com.co",
		Password: "ABC123@",
	}

	registerUser := &model.RegisteredUser{
		Email: "tigressa@gmail.sofka.com.co",
		Password: "ABC123@",
	}

	app.DataStore.On("ValidateUser",mock.Anything,mock.Anything).Return(user,nil)

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.ValidateUser(registerUser,user)

	assert.Nil(t,err)

	app.DataStore.MethodCalled("ValidateUser",mock.Anything)

}

func TestValidateUserIsWrongConecctionFailed(t *testing.T) {
	app:= test.CreateApp()

	user := &model.User{
		Id: "xxx1234",
		Name: "Tigre Maligno",
		Job: "developer",
		Description: "pruebas 2021",
		Cc: "123456789",
		Gender: "masculino",
		Email: "tigressa@gmail.sofka.com.co",
		Password: "ABC123@",
	}

	registerUser := &model.RegisteredUser{
		Email: "tigressa@gmail.sofka.com.co",
		Password: "ABC123@",
	}

	app.DataStore.On("ValidateUser",mock.Anything,mock.Anything).Return(nil,errors.New("UNREGISTERED_USER"))

	_, err := UserUseCase{
		datastore: app.DataStore,
	}.ValidateUser(registerUser,user)

	assert.NotNil(t,err,"UNREGISTERED_USER")

	app.DataStore.MethodCalled("ValidateUser",mock.Anything)

}

func TestUpdateUserIsOk(t *testing.T) {
	app:= test.CreateApp()

	user := &model.User{
		Id: "xxx1234",
		Name: "Tigre Maligno",
		Job: "developer",
		Description: "pruebas 2021",
		Cc: "123456789",
		Gender: "masculino",
		Email: "tigressa@gmail.sofka.com.co",
		Password: "ABC123@",
	}

	app.DataStore.On("UpdateUser",mock.Anything).Return(user,nil)


	 err := UserUseCase{
		datastore: app.DataStore,
	}.UpdateUser(user)

	assert.Nil(t,err)
	app.DataStore.MethodCalled("UpdateUser",mock.Anything)


}

func TestUpdateUserIsWrongConecctionFailed(t *testing.T) {
	app:= test.CreateApp()

	user := &model.User{
		Id: "xxx1234",
		Name: "Tigre Maligno",
		Job: "developer",
		Description: "pruebas 2021",
		Cc: "123456789",
		Gender: "masculino",
		Email: "tigressa@gmail.sofka.com.co",
		Password: "ABC123@",
	}

	app.DataStore.On("UpdateUser",mock.Anything).Return(nil,errors.New("CONNECTION_FAIL"))


	err := UserUseCase{
		datastore: app.DataStore,
	}.UpdateUser(user)

	assert.NotNil(t,err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("UpdateUser",mock.Anything)


}

func TestDeleteUserIsOk(t *testing.T) {
	app:= test.CreateApp()

	Id:= "xxx1234"

	app.DataStore.On("DeleteUser",mock.Anything).Return(nil)
	err := UserUseCase{
		datastore: app.DataStore,
	}.DeleteUser(Id)

	assert.Nil(t,err)
	app.DataStore.MethodCalled("DeleteUser",mock.Anything)


}

func TestDeleteUserIsWrongConecctionFailed(t *testing.T) {
	app:= test.CreateApp()

	Id:= "xxx1234"

	app.DataStore.On("DeleteUser",mock.Anything).Return(errors.New("CONNECTION_FAIL"))

	err := UserUseCase{
		datastore: app.DataStore,
	}.DeleteUser(Id)

	assert.NotNil(t,err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("DeleteUser",mock.Anything)


}

func TestRetrieverUserIsOk(t *testing.T) {
	app:= test.CreateApp()

	users := []model.User{{Id: "xxx1234",
		Name: "Tigre Maligno",
		Job: "developer",
		Description: "pruebas 2021",
		Cc: "123456789",
		Gender: "masculino",
		Email: "tigressa@gmail.sofka.com.co",
		Password: "ABC123@"},
	}

	app.DataStore.On("RetrieveUsers").Return(&users,nil)

	_,err := UserUseCase{
		datastore: app.DataStore,
	}.RetrieveUser()

	assert.Nil(t,err)
	app.DataStore.MethodCalled("RetrieveUsers")


}

func TestRetrieverUserIsWrongConecctionFailed(t *testing.T) {
	app:= test.CreateApp()

	app.DataStore.On("RetrieveUsers").Return(nil,errors.New("CONNECTION_FAIL"))

	_,err := UserUseCase{
		datastore: app.DataStore,
	}.RetrieveUser()

	assert.NotNil(t,err, "CONNECTION_FAIL")
	app.DataStore.MethodCalled("RetrieveUsers")


}