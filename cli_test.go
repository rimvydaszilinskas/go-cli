package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLI(t *testing.T) {
	cli := New("app")

	cmd, err := cli.AddCommand("run", "description", func(subcommands []string, flags []*Flag, flagValues map[string]interface{}) error {
		return nil
	})

	assert.NotNil(t, cmd)
	assert.Nil(t, err)

	err = cli.AddFlag(&Flag{
		Name: "type",
		Type: "int",
	})

	assert.Nil(t, err)

	command, err := cli.validateCommand("run")
	assert.Nil(t, err)
	assert.Equal(t, cmd, command)

	command, err = cli.validateCommand("notrun")
	assert.NotNil(t, err)
	assert.Nil(t, command)
}
