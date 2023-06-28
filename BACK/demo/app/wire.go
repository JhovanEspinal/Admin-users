//+build wireinject

package app

import "github.com/google/wire"

func CreateApp() *Application {

	wire.Build(
		DataStoreProvider,
		UserCasesProvider,
		AppProvider,
		)

	return new(Application)
}