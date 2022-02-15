package cli

import (
	"fmt"
	"strings"
)

// parseCommands finds all the commands in the call
func parseCommands(args []string) ([]string, error) {
	var commands []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--") || strings.HasPrefix(arg, "-") {
			break
		}
		commands = append(commands, arg)
	}

	comamndCount := len(commands)

	if comamndCount == 0 {
		return nil, fmt.Errorf("no command passed")
	}

	return commands, nil
}
