package user

import (
	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"flamingo.me/flamingo-app/src/user/domain"
	"flamingo.me/flamingo-app/src/user/infrastructure"
	graphqlinterface "flamingo.me/flamingo-app/src/user/interfaces/graphql"
)

// Module for user package
type Module struct{}

// Configure dingo
func (*Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(graphqlinterface.Service))
	injector.Bind(new(domain.UserService)).To(infrastructure.UserServiceImpl{})
}

// Depends defines the dependency to graphql.Module
func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(graphql.Module),
	}
}