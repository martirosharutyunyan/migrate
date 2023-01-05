package migrate

import (
	"log"
)

func Main(configFilePath string, args []string, cli bool) {
	InitConfig(configFilePath)
	InitDatabase()

	var cliArgument string

	if len(args) > 0 {
		cliArgument = args[0]
	} else {
		log.Fatal("run migrate help")
	}

	switch cliArgument {
	case "run":
		RunMigration(cli)
	case "revert":
		RevertMigration()
	case "generate":
		if len(args) != GENERATE_MIGRATION_ARGS_LEN {
			log.Fatal("Please provide the migration name")
		}
		GenerateMigration(args[1])
	case "help":
		Help()
	default:
		log.Fatal("run migrate help")
	}
}
