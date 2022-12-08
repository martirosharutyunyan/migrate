package main

import (
	"log"
	"os"

	"github.com/martirosharutyunyan/migrate/pkg"
)

func main() {
	var cliArgument string

	if len(os.Args) > 1 {
		cliArgument = os.Args[1]
	} else {
		log.Fatal("run go help")
	}

	switch cliArgument {
	case "run":
		migrate.RunMigration()
	case "revert":
		migrate.RevertMigration()
	case "generate":
		if len(os.Args) != migrate.GENERATE_MIGRATION_ARGS_LEN {
			log.Fatal("Please provide the migration name")
		}
		migrate.GenerateMigration(os.Args[2])
	case "help":
		migrate.Help()
	default:
		log.Fatal("run migrate help")
	}
}
