package migrate

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func init() {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v",
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
