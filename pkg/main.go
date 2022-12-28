package migrate

import (
	"log"
	"os"
)

func Main(configFilePath string) {
	InitConfig(configFilePath)
	InitDatabase()

	var cliArgument string

	if len(os.Args) > 1 {
		cliArgument = os.Args[1]
	} else {
		log.Fatal("run go help")
	}

	switch cliArgument {
	case "run":
		RunMigration()
	case "revert":
		RevertMigration()
	case "generate":
		if len(os.Args) != GENERATE_MIGRATION_ARGS_LEN {
			log.Fatal("Please provide the migration name")
		}
		GenerateMigration(os.Args[2])
	case "help":
		Help()
	default:
		log.Fatal("run migrate help")
	}
}
