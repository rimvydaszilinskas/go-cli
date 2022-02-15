package cli

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-wordwrap"
)

type CLI struct {
	Name     string
	Commands []*command
	Flags    []*Flag

	LineLength uint
	FlagValues map[string]interface{}

	parsedFlags []*Flag
	Description string
}

func New(name string, commands ...*command) *CLI {
	return &CLI{
		Name:       name,
		LineLength: 80,
		Commands:   commands,
	}
}

func (cli *CLI) AddCommand(name string, description string, exec runCommand) (*command, error) {
	for _, cmd := range cli.Commands {
		if cmd.Name == name {
			return nil, fmt.Errorf("command %s already exists", name)
		}
	}
	cmd := &command{
		Name:        name,
		Description: description,
		Command:     exec,
	}
	cli.Commands = append(cli.Commands, cmd)
	return cmd, nil
}

func (cli *CLI) AddFlag(flag *Flag) error {
	for _, f := range cli.Flags {
		if f.Name == flag.Name {
			return fmt.Errorf("flag already exists")
		}
	}
	cli.Flags = append(cli.Flags, flag)
	return nil
}

func (cli *CLI) Run() {
	if err := cli.Execute(); err != nil {
		fatal(err)
	}
}

func (cli *CLI) Execute() error {
	commandStrings, err := parseCommands(os.Args[1:])

	if err != nil {
		return err
	}

	if commandStrings[0] == "help" {
		cli.DisplayAllCommands()
		return nil
	}

	comamnd, err := cli.validateCommand(commandStrings[0])

	if err != nil {
		return err
	}

	if len(commandStrings) > 1 {
		if commandStrings[1] == "help" {
			cli.DisplayDescription(comamnd)
			return nil
		}
	}

	if err := cli.parseFlags(comamnd, len(commandStrings)); err != nil {
		return err
	}

	return comamnd.Command(commandStrings[1:], cli.parsedFlags, cli.FlagValues)
}

func (cli *CLI) validateCommand(commandName string) (*command, error) {
	for _, command := range cli.Commands {
		if command.Name == commandName {
			return command, nil
		}
	}
	return nil, fmt.Errorf("invalid command '%s'. Use 'help' to display all commands", commandName)
}

func (cli *CLI) parseFlags(command *command, commandCount int) error {
	flags := command.Flags

	for _, flag := range cli.Flags {
		if !flagAlreadyExists(flags, flag) {
			flags = append(flags, flag)
		}
	}

	values := make(map[string]interface{})

	for _, flag := range flags {
		if err := flag.Init(); err != nil {
			return fmt.Errorf("error initializing flag %s - %w", flag.Name, err)
		}
	}

	flag.CommandLine.Parse(os.Args[1+commandCount:])

	for _, flag := range flags {
		flag.Validate()
		if flag.IsBool() {
			values[flag.Name] = flag.GetBool()
		} else if flag.IsInt() {
			values[flag.Name] = flag.GetInt()
		} else if flag.IsString() {
			values[flag.Name] = flag.GetString()
		} else if flag.IsFloat() {
			values[flag.Name] = flag.GetFloat()
		}
	}

	cli.FlagValues = values
	cli.parsedFlags = flags
	return nil
}

func (cli *CLI) DisplayDescription(command *command) {
	tmpl := fmt.Sprintf("%s command line\n%s", cli.Name, cli.LineBreak())
	tmpl += fmt.Sprintf("Command %s\n", command.Name)

	if len(command.Description) == 0 {
		tmpl += "No description provided"
	} else {
		tmpl += "- " + command.Description + "\n"
	}

	if len(command.Flags) != 0 {
		tmpl += "\nThe following flags can be set:\n\n"

		for _, flag := range command.Flags {
			tmpl += flagDescription(flag)
		}
	}

	tmpl += cli.LineBreak()
	tmpl += DescriptionFooter

	fmt.Println(wordwrap.WrapString(tmpl, cli.LineLength))
}

func (cli *CLI) DisplayAllCommands() {
	tmpl := fmt.Sprintf("%s command line\n%sThe following commands are available:\n\n", cli.Name, cli.LineBreak())

	for _, cmd := range cli.Commands {
		tmpl += fmt.Sprintf("%s - %s\n", cmd.Name, cmd.Description)
	}

	if len(cli.Flags) != 0 {
		tmpl += "\nThe following global flags can be set:\n\n"

		for _, flag := range cli.Flags {
			tmpl += flagDescription(flag)
		}
	}

	tmpl += cli.LineBreak()
	tmpl += DescriptionFooter

	fmt.Println(wordwrap.WrapString(tmpl, cli.LineLength))
}

func (cli *CLI) LineBreak() string {
	return fmt.Sprintf("%s\n", strings.Repeat("-", int(cli.LineLength)))
}

func flagDescription(flag *Flag) string {
	if flag.Default != nil {
		return fmt.Sprintf("--%s (type: %s default: %s) = %s\n", flag.Name, flag.Type, flag.Default, flag.Description)
	}
	return fmt.Sprintf("--%s (type: %s) = %s\n", flag.Name, flag.Type, flag.Description)
}

func flagAlreadyExists(flags []*Flag, flag *Flag) bool {
	for _, f := range flags {
		if f.Name == flag.Name {
			return true
		}
	}
	return false
}
