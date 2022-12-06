package migrate

import "fmt"

func Help() {
	fmt.Println(`
Commands:
	run: Synchronizes Database
	revert: Reverts last migration
	generate: Generates new migration with argument name
	help: Help
	`)
}
