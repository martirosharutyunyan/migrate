package migrate

import (
	"fmt"
	"log"
	"os"
	"time"
)

func GenerateMigration(name string) {
	timestamp := time.Now().Unix()
	migrationUpFile, migrationUpFileErr := os.Create(fmt.Sprintf("%v/%v-%v.up.sql", config["migrations"], timestamp, name))
	migrationDownFile, migrationDownFileErr := os.Create(fmt.Sprintf("%v/%v-%v.down.sql", config["migrations"], timestamp, name))
	defer migrationUpFile.Close()
	defer migrationDownFile.Close()

	if migrationUpFileErr != nil {
		log.Fatal(migrationUpFileErr)
	}

	if migrationDownFileErr != nil {
		log.Fatal(migrationDownFileErr)
	}

}
