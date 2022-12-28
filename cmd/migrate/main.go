package main

import (
	"github.com/martirosharutyunyan/migrate/pkg"
)

func main() {
	migrate.Main("./dbconfig.json")
}
