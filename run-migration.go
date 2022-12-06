package migrate

import (
	"log"
	"os"
)

type MigrationEntity struct {
	Name string
}

func RunMigration() {
	folderMigrations, directoryErr := os.ReadDir(config["migrations"].(string))

	if directoryErr != nil {
		log.Fatal(directoryErr)
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
		db.Exec(migrationSql)
		db.Exec("INSERT INTO migrations (name) VALUES (?)", migrationName)
	}

	log.Println("Migrations synced to database")
}
