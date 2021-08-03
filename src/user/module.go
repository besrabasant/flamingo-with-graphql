package user

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo-app/src/user/commands"
	"flamingo.me/flamingo-app/src/user/domain"
	"flamingo.me/flamingo-app/src/user/infrastructure"
	graphqlinterface "flamingo.me/flamingo-app/src/user/interfaces/graphql"
	"flamingo.me/graphql"
	"github.com/spf13/cobra"
)

// Module for user package
type Module struct{}

// Configure dingo
func (*Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(graphqlinterface.Service))
	injector.Bind(new(domain.UserService)).To(infrastructure.UserServiceImpl{})
	injector.BindMulti(new(cobra.Command)).ToInstance(commands.MigrateSchemaCommand())
}

// Depends defines the dependency to graphql.Module
func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(graphql.Module),
	}
}
