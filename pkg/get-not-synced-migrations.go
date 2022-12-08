package migrate

import (
	"log"
)

func GetNotSyncedMigrations(folderMigrations []string, migrationEntities []MigrationEntity) []string {
	var notSyncedMigrations []string
	for _, folderMigrationName := range folderMigrations {
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

		for _, folderMigrationName := range folderMigrations {
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
