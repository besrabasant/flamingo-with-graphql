package todo

import (
	"flamingo.me/dingo"
	"flamingo.me/graphql"
	"flamingo.me/flamingo-app/src/todo/domain"
	"flamingo.me/flamingo-app/src/user"
)

type Service struct{}

func (*Service) Schema() []byte {
	//language=graphql
	return []byte(`
	type Todo {
		id: ID!
		task: String!
		done: Boolean!
	}
	extend type User {
		todos: [Todo]
	}
	extend type Mutation {
		TodoAdd(user: ID!, task: String!): Todo
		TodoDone(todo: ID!, done: Boolean!): Todo
	}	
	`)
}

func (*Service) Types(types *graphql.Types) {
	types.Map("Todo", domain.Todo{})
	types.Resolve("User", "todos", UserResolver{}, "Todos")
	types.Resolve("Mutation", "TodoAdd", MutationResolver{}, "TodoAdd")
	types.Resolve("Mutation", "TodoDone", MutationResolver{}, "TodoDone")
}

type Module struct{}

func (Module) Configure(injector *dingo.Injector) {
	injector.BindMulti(new(graphql.Service)).To(new(Service))
}

func (*Module) Depends() []dingo.Module {
	return []dingo.Module{
		new(user.Module),
	}
}
