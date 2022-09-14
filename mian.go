package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type MigrationEntity struct {
	gorm.Model
	Name string
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

	cliArgument := os.Args[1]

	switch cliArgument {
	case "--sync":
		MigrationSync(configJson)
		fmt.Println("Migrations synced to database")
	case "--generate":
		if len(os.Args) != 3 {
			log.Fatal("Please provide the migration name")
		}
		GenerateMigration(os.Args[2], configJson)
	}
}

func MigrationSync(configJson map[string]any) {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", configJson["host"], configJson["user"], configJson["password"], configJson["dbname"], configJson["port"])

	db, connectionErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if connectionErr != nil {
		fmt.Println(connectionErr)
	}

	autoMigrateErr := db.AutoMigrate(&MigrationEntity{})
	if autoMigrateErr != nil {
		fmt.Println(autoMigrateErr)
	}

	folderMigrations, directoryErr := os.ReadDir(configJson["migrations"].(string))

	if directoryErr != nil {
		log.Fatal(directoryErr)
	}

	var migrationEntities []MigrationEntity
	db.Model(&MigrationEntity{}).Scan(&migrationEntities)

	notSyncedMigrations := GetNotSyncedMigrations(folderMigrations, migrationEntities)

	for _, migrationName := range notSyncedMigrations {
		migrationSqlBuffer, readFileErr := os.ReadFile(configJson["migrations"].(string) + "/" + migrationName)

		if readFileErr != nil {
			log.Fatal(readFileErr)
		}

		migrationSql := string(migrationSqlBuffer)
		db.Exec(migrationSql)
		db.Create(&MigrationEntity{Name: migrationName})
	}
}

func GenerateMigration(name string, configJson map[string]any) {
	timestamp := time.Now().Unix()
	file, err := os.Create(fmt.Sprintf("%v/%v-%v.sql", configJson["migrations"], timestamp, name))
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func GetNotSyncedMigrations(folderMigrations []os.DirEntry, migrationEntities []MigrationEntity) []string {
	var notSyncedMigrations []string
	for _, folderMigration := range folderMigrations {
		existMigration := false

		for _, migrationEntity := range migrationEntities {
			if folderMigration.Name() == migrationEntity.Name {
				existMigration = true
			}
		}

		if !existMigration {
			notSyncedMigrations = append(notSyncedMigrations, folderMigration.Name())
		}

	}

	var deletedMigrations []string
	for _, migrationEntity := range migrationEntities {
		existMigration := false

		for _, folderMigration := range folderMigrations {
			if folderMigration.Name() == migrationEntity.Name {
				existMigration = true
			}
		}

		if !existMigration {
			deletedMigrations = append(deletedMigrations, migrationEntity.Name)
		}
	}

	if len(deletedMigrations) != 0 {
		fmt.Println("From folder deleted migrations :", deletedMigrations)
	}

	return notSyncedMigrations
}
