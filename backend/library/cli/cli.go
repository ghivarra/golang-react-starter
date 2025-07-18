package cli

import (
	"backend/library/migration"
	"fmt"
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

	fmt.Println(commands)
}
