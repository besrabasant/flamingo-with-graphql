package graphql

import (
	"flamingo.me/flamingo-app/src/user/domain"
	"flamingo.me/flamingo-app/src/utils"
	"flamingo.me/graphql"
)

//go:generate go run github.com/go-bindata/go-bindata/v3/go-bindata -nometadata -o fs.go -pkg graphql schema.graphql

// Service service for graphql
type Service struct{}

// Schema defines the graphql schema
func (*Service) Schema() []byte {
	return utils.MustAsset("schema.graphql")
}

// Types set up the GraphQL to Go Type mappings
func (*Service) Types(types *graphql.Types) {
	types.Map("User", domain.User{})
	types.Map("User_Attributes", domain.Attributes{})
	types.Resolve("Query", "User", UserQueryResolver{}, "User")
}
