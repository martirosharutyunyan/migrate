package main

import (
	"github.com/martirosharutyunyan/migrate"
	"log"
	"os"
)

func main() {
	var cliArgument string

	if len(os.Args) > 1 {
		cliArgument = os.Args[1]
	} else {
		log.Fatal("run go help")
	}

	switch cliArgument {
	case "--run":
		migrate.RunMigration()
	case "--revert":
		migrate.RevertMigration()
	case "--help":
	case "--generate":
		if len(os.Args) != 3 {
			log.Fatal("Please provide the migration name")
		}
		migrate.GenerateMigration(os.Args[2])
	case "help":
		migrate.Help()
	default:
		log.Fatal("run go help")
	}
}
