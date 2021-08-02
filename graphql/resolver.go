// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

//+build !graphql

package graphql

import (
	"context"

	"flamingo.me/flamingo-app/src/todo"
	"flamingo.me/flamingo-app/src/todo/domain"
	domain1 "flamingo.me/flamingo-app/src/user/domain"
	graphql2 "flamingo.me/flamingo-app/src/user/interfaces/graphql"
	graphql1 "flamingo.me/graphql"
)

var _ ResolverRoot = new(rootResolver)

type rootResolver struct {
	rootResolverMutation *rootResolverMutation
	rootResolverQuery    *rootResolverQuery
	rootResolverUser     *rootResolverUser
}

func (r *rootResolver) Inject(
	rootResolverMutation *rootResolverMutation,
	rootResolverQuery *rootResolverQuery,
	rootResolverUser *rootResolverUser,
) {
	r.rootResolverMutation = rootResolverMutation
	r.rootResolverQuery = rootResolverQuery
	r.rootResolverUser = rootResolverUser
}

func (r *rootResolver) Mutation() MutationResolver {
	return r.rootResolverMutation
}
func (r *rootResolver) Query() QueryResolver {
	return r.rootResolverQuery
}
func (r *rootResolver) User() UserResolver {
	return r.rootResolverUser
}

type rootResolverMutation struct {
	resolveFlamingo func(ctx context.Context) (*string, error)
	resolveTodoAdd  func(ctx context.Context, user string, task string) (*domain.Todo, error)
	resolveTodoDone func(ctx context.Context, todo string, done bool) (*domain.Todo, error)
}

func (r *rootResolverMutation) Inject(
	mutationFlamingo *graphql1.FlamingoQueryResolver,
	mutationTodoAdd *todo.MutationResolver,
	mutationTodoDone *todo.MutationResolver,
) {
	r.resolveFlamingo = mutationFlamingo.Flamingo
	r.resolveTodoAdd = mutationTodoAdd.TodoAdd
	r.resolveTodoDone = mutationTodoDone.TodoDone
}

func (r *rootResolverMutation) Flamingo(ctx context.Context) (*string, error) {
	return r.resolveFlamingo(ctx)
}
func (r *rootResolverMutation) TodoAdd(ctx context.Context, user string, task string) (*domain.Todo, error) {
	return r.resolveTodoAdd(ctx, user, task)
}
func (r *rootResolverMutation) TodoDone(ctx context.Context, todo string, done bool) (*domain.Todo, error) {
	return r.resolveTodoDone(ctx, todo, done)
}

type rootResolverQuery struct {
	resolveFlamingo func(ctx context.Context) (*string, error)
	resolveUser     func(ctx context.Context, id string) (*domain1.User, error)
}

func (r *rootResolverQuery) Inject(
	queryFlamingo *graphql1.FlamingoQueryResolver,
	queryUser *graphql2.UserQueryResolver,
) {
	r.resolveFlamingo = queryFlamingo.Flamingo
	r.resolveUser = queryUser.User
}

func (r *rootResolverQuery) Flamingo(ctx context.Context) (*string, error) {
	return r.resolveFlamingo(ctx)
}
func (r *rootResolverQuery) User(ctx context.Context, id string) (*domain1.User, error) {
	return r.resolveUser(ctx, id)
}

type rootResolverUser struct {
	resolveTodos func(ctx context.Context, obj *domain1.User) ([]*domain.Todo, error)
}

func (r *rootResolverUser) Inject(
	userTodos *todo.UserResolver,
) {
	r.resolveTodos = userTodos.Todos
}

func (r *rootResolverUser) Todos(ctx context.Context, obj *domain1.User) ([]*domain.Todo, error) {
	return r.resolveTodos(ctx, obj)
}

func direct(root *rootResolver) map[string]interface{} {
	return map[string]interface{}{
		"Mutation.Flamingo": root.Mutation().Flamingo,
		"Mutation.TodoAdd":  root.Mutation().TodoAdd,
		"Mutation.TodoDone": root.Mutation().TodoDone,
		"Query.Flamingo":    root.Query().Flamingo,
		"Query.User":        root.Query().User,
		"User.Todos":        root.User().Todos,
	}
}
