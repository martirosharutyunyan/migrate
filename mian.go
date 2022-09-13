package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type MigrationEntity struct {
	gorm.Model
	name string
}

func main() {
	content, err := os.ReadFile("./dbconfig.json")
	if err != nil {
		log.Fatal(err)
	}

	var configJson = make(map[string]any)
	marshalErr := json.Unmarshal(content, &configJson)

	if marshalErr != nil {
		log.Fatal(marshalErr)
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", configJson["host"], configJson["user"], configJson["password"], configJson["dbname"], configJson["port"])

	db, connectionErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if connectionErr != nil {
		fmt.Println(connectionErr)
	}

	db.AutoMigrate(&MigrationEntity{})

	fmt.Println("Connected to database")

	folderMigrations, directoryErr := os.ReadDir(configJson["migrations"].(string))

	if directoryErr != nil {
		log.Fatal(directoryErr)
	}
	var migrationNames []string
	for _, migration := range folderMigrations {
		migrationNames = append(migrationNames, migration.Name())
	}

	var migrationEntities []MigrationEntity
	db.Model(&MigrationEntity{}).Scan(migrationEntities)

	fmt.Println(migrationEntities)

	var migrationDifferences []string
	for _, migrationName := range migrationNames {
		for _, entity := range migrationEntities {
			if migrationName != entity.name {
				migrationDifferences = append(migrationDifferences, migrationName)
			}
		}
	}

	for _, migrationName := range migrationDifferences {
		migrationSql, readFileErr := os.ReadFile(configJson["migrations"].(string) + migrationName)

		if readFileErr != nil {
			log.Fatal(readFileErr)
		}

		fmt.Println(migrationSql)
	}

}
