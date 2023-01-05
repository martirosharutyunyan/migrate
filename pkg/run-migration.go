package migrate

import (
	"log"
	"os"
	"strings"
)

type MigrationEntity struct {
	Name string
}

func RunMigration(cli bool) {
	migrationFolderPath := GetMigrationFolderPath()

	if !cli {
		migrationFolderPath = "../../" + migrationFolderPath
	}

	folderMigrationDirEntries, directoryErr := os.ReadDir(migrationFolderPath)

	if directoryErr != nil {
		log.Fatal(directoryErr)
	}

	var folderMigrations []string
	for _, folderMigrationDirEntry := range folderMigrationDirEntries {
		if strings.Contains(folderMigrationDirEntry.Name(), ".up.sql") {
			folderMigrations = append(folderMigrations, folderMigrationDirEntry.Name())
		}
	}

	var migrationEntities []MigrationEntity
	rows, queryErr := db.Query("SELECT name FROM migrations")

	if queryErr != nil {
		log.Fatal(queryErr)
	}
	defer rows.Close()

	for rows.Next() {
		var migrationEntity MigrationEntity
		scanErr := rows.Scan(&migrationEntity.Name)

		if scanErr != nil {
			log.Fatal(scanErr)
		}
		migrationEntities = append(migrationEntities, migrationEntity)
	}

	notSyncedMigrations := GetNotSyncedMigrations(folderMigrations, migrationEntities)

	for _, migrationName := range notSyncedMigrations {
		migrationSQLBuffer, readFileErr := os.ReadFile(migrationFolderPath + "/" + migrationName)

		if readFileErr != nil {
			log.Fatal(readFileErr)
		}

		migrationSQL := string(migrationSQLBuffer)

		_, migrationErr := db.Exec(migrationSQL)
		if migrationErr != nil {
			log.Fatal(migrationErr)
		}

		_, insertErr := db.Exec("INSERT INTO migrations (name) VALUES ($1)", migrationName)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
	}

	log.Println("Migrations synced to database: ", notSyncedMigrations)
}
