package main

import (
	"github.com/rimvydaszilinskas/go-cli"
)

func main() {
	app := cli.New("cli")

	app.AddCommand("hello", "say hello to the world", func(subcommands []string, flags []*cli.Flag, flagValues map[string]interface{}) error {
		print("hello")
		return nil
	})

	app.Execute()
}
