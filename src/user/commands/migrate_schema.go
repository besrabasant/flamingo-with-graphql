package commands

import (
	"fmt"

	"flamingo.me/flamingo-app/src/database"
	userInfrastructure "flamingo.me/flamingo-app/src/user/infrastructure"
	"github.com/spf13/cobra"
)

func MigrateSchemaCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrates database schema.",
		Run:   MigrateSchemaCommandHandler,
	}

	return cmd
}

func MigrateSchemaCommandHandler(cmd *cobra.Command, args []string) {
	fmt.Println("Migrating user schema...")
	database.Connection().AutoMigrate(&userInfrastructure.User{})

	if database.Connection().Migrator().HasColumn(&userInfrastructure.User{}, "email") {
		database.Connection().Migrator().DropColumn(&userInfrastructure.User{}, "email")
	}
}
