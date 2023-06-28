package app

import (
	"demo/domain/useCases"
	"demo/infra/dataStore"
	"github.com/google/wire"
)

var DataStoreProvider = wire.NewSet(dataStore.NewMongoGatewayImpl)
var UserCasesProvider = wire.NewSet(useCases.NewUserUseCase)
var AppProvider = wire.NewSet(NewApplication)
