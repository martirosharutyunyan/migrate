package migrate

import (
	"log"
	"os"
	"strings"
)

func GetNotSyncedMigrations(folderMigrations []os.DirEntry, migrationEntities []MigrationEntity) []string {
	var notSyncedMigrations []string
	for _, folderMigration := range folderMigrations {
		folderMigrationName := folderMigration.Name()
		if strings.Contains(folderMigrationName, ".down.sql") {
			continue
		}

		existMigration := false

		for _, migrationEntity := range migrationEntities {
			if folderMigrationName == migrationEntity.Name {
				existMigration = true
			}
		}

		if !existMigration {
			notSyncedMigrations = append(notSyncedMigrations, folderMigrationName)
		}
	}

	var deletedMigrations []string
	for _, migrationEntity := range migrationEntities {
		existMigration := false

		for _, folderMigration := range folderMigrations {
			folderMigrationName := folderMigration.Name()
			if strings.Contains(folderMigrationName, ".down.sql") {
				continue
			}

			if folderMigrationName == migrationEntity.Name {
				existMigration = true
			}
		}

		if !existMigration {
			deletedMigrations = append(deletedMigrations, migrationEntity.Name)
		}
	}

	if len(deletedMigrations) != 0 {
		log.Println("From folder deleted migrations :", deletedMigrations)
	}

	return notSyncedMigrations
}
