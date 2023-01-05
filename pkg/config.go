package migrate

import (
	"encoding/json"
	"log"
	"os"
)

var config map[string]any

func InitConfig(configFilePath string) {
	content, err := os.ReadFile(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	marshalErr := json.Unmarshal(content, &config)

	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
}

func GetMigrationFolderPath() string {
	return config["migrations"].(string)
}
