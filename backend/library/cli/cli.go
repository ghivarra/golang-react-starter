package cli

import (
	"backend/library/cli/maker"
	"backend/library/migration"
	"os"
	"strings"
)

// Run CLI Commands
func Run(commands string) {

	// slice commands
	commandArray := strings.Split(commands, ":")

	// check if db
	if commandArray[0] == "db" {
		migration.Run(commandArray[1])
		return
	}

	if commandArray[0] == "make" && len(os.Args) > 1 {
		tableName := os.Args[2]
		maker.BuildMigration(tableName)
	}

	// fmt.Println(commands)
}
