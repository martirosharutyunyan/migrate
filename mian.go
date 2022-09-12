package main

import (
	"database/sql"
	"fmt"
	_ "github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	dsn := "postgresql://postgres:postgres@localhost:5432/learning"

	_, err := sql.Open("pg", dsn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Connected")
}
