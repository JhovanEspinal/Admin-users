package test

import (
	"demo/domain"
	"github.com/google/wire"
)

var DbGateweyProvider = wire.NewSet(NewDbGateway, wire.Bind(new(domain.DatabaseGateway),new(DbGateway)))

var TestApplicacion = wire.NewSet(NewApplication)
