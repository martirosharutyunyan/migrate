package main

import (
	"github.com/martirosharutyunyan/migrate/pkg"
	"os"
)

func main() {
	migrate.Main("./dbconfig.json", os.Args, true)
}
