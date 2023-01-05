package migrate

import (
	"log"
	"os"
)

func RevertMigration() {
	folderMigrations, directoryErr := os.ReadDir(GetMigrationFolderPath())

	if directoryErr != nil {
		log.Fatal(directoryErr)
	}

	migrationDownName := folderMigrations[len(folderMigrations)-2].Name()
	migrationUpName := folderMigrations[len(folderMigrations)-1].Name()

	migrationSQLBuffer, readFileErr := os.ReadFile(GetMigrationFolderPath() + "/" + migrationDownName)

	if readFileErr != nil {
		log.Fatal(readFileErr)
	}

	migrationSql := string(migrationSQLBuffer)

	_, migrationErr := db.Exec(migrationSql)
	if migrationErr != nil {
		log.Fatal(migrationErr)
	}

	_, dropErr := db.Exec("DELETE FROM migrations WHERE name=$1", migrationUpName)
	if dropErr != nil {
		log.Fatal(dropErr)
	}

	log.Printf("Migration %v reverted", migrationUpName)
}
