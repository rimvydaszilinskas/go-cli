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

func NewCommand(name, description string, exec runCommand) (*command, error) {
	if name != strings.ToLower(name) || len(name) == 0 {
		return nil, fmt.Errorf("invalid command name: '%s'", name)
	}
	return &command{
		Name:        name,
		Description: description,
		Command:     exec,
	}, nil
}

func (cmd *command) AddFlag(flag *Flag) error {
	for _, f := range cmd.Flags {
		if f.Name == flag.Name {
			return fmt.Errorf("flag already exists")
		}
	}

	cmd.Flags = append(cmd.Flags, flag)
	return nil
}
