package migrate

import (
	"database/sql"
	"fmt"
	"log"

	// postgresql driver
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDatabase() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		config["host"],
		config["user"],
		config["password"],
		config["dbname"],
		config["port"])

	var err error
	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS migrations (name varchar)")

	if err != nil {
		log.Fatal(err)
	}
}
