package cli

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitFlag(t *testing.T) {
	flag := Flag{
		Name: "git",
		Type: "int",
	}

	err := flag.Init()

	assert.Nil(t, err)
	assert.Equal(t, -1, flag.Default)
}

func TestFlagValidation(t *testing.T) {
	validationFunction := func(val interface{}) error {
		if val.(int) == 100 {
			return fmt.Errorf("100")
		}
		return nil
	}

	flag := Flag{
		Name:       "git",
		Type:       "int",
		Value:      123,
		Validation: &validationFunction,
	}

	err := flag.Validate()
	assert.Nil(t, err)

	flag.Value = 100
	err = flag.Validate()
	assert.NotNil(t, err)

	flag.Validation = nil
	flag.Value = "123a"

	err = flag.Validate()
	assert.NotNil(t, err)
}

func TestIncorrectTypes(t *testing.T) {
	flag := Flag{
		Name:    "git",
		Type:    "int",
		Default: "commit",
	}

	err := flag.Init()

	assert.NotNil(t, err)
}
