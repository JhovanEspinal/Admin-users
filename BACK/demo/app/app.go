package app

import (
	"demo/domain"
	"demo/domain/useCases"
	"github.com/gorilla/mux"
)

type Application struct {

	Router *mux.Router
	datastore domain.DatabaseGateway
	userUseCases useCases.UserUseCase
}

func NewApplication(datastore domain.DatabaseGateway, userUseCase useCases.UserUseCase) *Application{

	return &Application{
		datastore: datastore,
		userUseCases: userUseCase,
	}

}

