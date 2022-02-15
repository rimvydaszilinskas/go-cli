package cli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommands(t *testing.T) {
	args := []string{
		"git", "clone", "--do", "nothing",
	}

	commands, err := parseCommands(args)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(commands))
	assert.Equal(t, "git", commands[0])
	assert.Equal(t, "clone", commands[1])

	args = []string{
		"git", "--do", "nothing",
	}

	commands, err = parseCommands(args)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(commands))
	assert.Equal(t, "git", commands[0])
}

func TestNoCommandsParsed(t *testing.T) {
	args := []string{
		"--do", "something",
	}
	commands, err := parseCommands(args)
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(commands))
}
