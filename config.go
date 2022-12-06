package migrate

import (
	"encoding/json"
	"log"
	"os"
)

var config map[string]any

func init() {
	content, err := os.ReadFile("./dbconfig.json")
	if err != nil {
		log.Fatal(err)
	}

	marshalErr := json.Unmarshal(content, &config)

	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
}
