package migrations

import (
	"fmt"
)

type CreateTodosTable struct{}

func (*CreateTodosTable) up() {
	fmt.Println("CreateTodosTable migrate up")
}

func (*CreateTodosTable) down() {
	fmt.Println("CreateTodosTable migrate down")
}
