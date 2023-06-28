package app

import (
	"demo/domain/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func (app Application) SaveUser(responseWriter http.ResponseWriter, request *http.Request){

	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(responseWriter,err.Error(),http.StatusBadRequest)
		return
	}

	userSaved, err := app.userUseCases.SaveUser(&user)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(userSaved)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("content-type","application/json")
	responseWriter.Write(data)
}

func (app Application) ValidateUser(responseWriter http.ResponseWriter, request *http.Request){

	var userData model.RegisteredUser
	var user model.User
	err := json.NewDecoder(request.Body).Decode(&userData)
	if err != nil {
		http.Error(responseWriter,err.Error(),http.StatusBadRequest)
		return
	}

	ValidatedUser, err := app.userUseCases.ValidateUser(&userData,&user)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(ValidatedUser)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("content-type","application/json")
	responseWriter.Write(data)
}

func (app Application) RetrieveUsers(responseWriter http.ResponseWriter, request *http.Request) {

	users, err := app.userUseCases.RetrieveUser()
	if err != nil{
		 http.Error(responseWriter,err.Error(),http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(users)
	if err != nil {
		http.Error(responseWriter,err.Error(),http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type","application/json")
	responseWriter.Write(data)
}

func (app Application) DeleteUser(responseWriter http.ResponseWriter, request *http.Request) {

	id := mux.Vars(request)["id"]

	err := app.userUseCases.DeleteUser(id)
	if err != nil{
		http.Error(responseWriter,err.Error(),http.StatusInternalServerError)
		return
	}
}

func (app Application) UpdateUser(responseWriter http.ResponseWriter, request *http.Request) {

	var user model.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Id == "" {
		http.Error(responseWriter, "USER_NOT_FOUND", http.StatusBadRequest)
		return
	}

	err = app.userUseCases.UpdateUser(&user)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
}
