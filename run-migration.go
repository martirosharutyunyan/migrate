package migrate

import (
	"log"
	"os"
	"strings"
)

type MigrationEntity struct {
	Name string
}

func RunMigration() {
	folderMigrationDirEntries, directoryErr := os.ReadDir(config["migrations"].(string))

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
		migrationSqlBuffer, readFileErr := os.ReadFile(config["migrations"].(string) + "/" + migrationName)

		if readFileErr != nil {
			log.Fatal(readFileErr)
		}

		migrationSql := string(migrationSqlBuffer)

		_, migrationErr := db.Exec(migrationSql)
		if migrationErr != nil {
			log.Fatal(migrationErr)
		}

		_, insertErr := db.Exec("INSERT INTO migrations (name) VALUES ($1)", migrationName)
		if insertErr != nil {
			log.Fatal(insertErr)
		}
	}

	log.Println("Migrations synced to database")
}
