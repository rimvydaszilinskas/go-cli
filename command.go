package cli

import (
	"fmt"
	"strings"
)

type runCommand = func(subcommands []string, flags []*Flag, flagValues map[string]interface{}) error

type command struct {
	Name        string
	Command     runCommand
	Description string
	Flags       []*Flag
}

func NewCommand(name, description string, exec runCommand) *command {
	if name != strings.ToLower(name) || len(name) == 0 {
		panic(fmt.Sprintf("invalid command name: '%s'", name))
	}
	return &command{
		Name:        name,
		Description: description,
		Command:     exec,
	}
}

func (cmd *command) AddFlag(flag *Flag) {
	cmd.Flags = append(cmd.Flags, flag)
}
