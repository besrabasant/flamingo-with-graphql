package main

//go:generate rm -f graphql/generated.go
//go:generate go run -tags graphql main.go graphql

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/gotemplate"
	"flamingo.me/flamingo/v3/core/requestlogger"
	"flamingo.me/flamingo/v3/core/zap"

	projectGraphql "flamingo.me/flamingo-app/graphql"
	"flamingo.me/flamingo-app/src/database"
	"flamingo.me/flamingo-app/src/helloworld"
	"flamingo.me/flamingo-app/src/todo"
)

func main() {

	flamingo.App([]dingo.Module{
		new(zap.Module),
		new(requestlogger.Module),
		new(gotemplate.Module),
		new(database.Module),
		new(projectGraphql.Module),
		new(helloworld.Module),
		new(todo.Module),
	})
}
