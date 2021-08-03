package database

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo-app/src/database/infrastructure"
	"flamingo.me/flamingo-app/src/database/interfaces"
)

type Module struct{}

func (*Module) Configure(injector *dingo.Injector) {
	var dbServiceInstance = new(infrastructure.DBService)

	dbServiceInstance.SetConnection(Connection())

	injector.Bind(new(interfaces.DB)).ToInstance(dbServiceInstance)
}
