package migrate

import (
	"log"
	"os"
)

func RevertMigration() {
	folderMigrations, directoryErr := os.ReadDir(config["migrations"].(string))

	if directoryErr != nil {
		log.Fatal(directoryErr)
	}

	for _, migration := range folderMigrations {
		log.Println(migration.Name())
	}

}
