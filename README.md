# CLI

## Features

- Easy subcommand parsing
- Easy flag definition and validation
- Help message formatting

## Call structure

```sh
app <comamnd_name> <subcommands> <flags>
```

Example:

```
app hello once -count=1
```

## Help

You can always run the application with command `help` which will default to displaying al;l the information about the app:

```sh
app help
```

Produces:

```
cli command line
--------------------------------------------------------------------------------
The following commands are available:

hello - say hello to the world
--------------------------------------------------------------------------------
Powered by github.com/rimvydaszilinskas/go-cli
```

You can also get help for a specific command.

```sh
app hello help
```

Produces:

```
cli command line
--------------------------------------------------------------------------------
Command hello
- say hello to the world
--------------------------------------------------------------------------------
Powered by github.com/rimvydaszilinskas/go-cli
```

## Installation

```sh
go get github.com/rimvydaszilinskas/go-cli
```

## Usage

Here is a basic usage:

```go
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

	app.Run()
}
```

### Specifying flags for all commands

```go
app.AddFlag(&Flag{
    Name: "count",
    Type: "int",
    Description: "How many times should it print hello",
    Default: 1,
})
```

### Specifying flags for a command

```go
cmd := app.AddCommand("hello", "say hello to the world", func(subcommands []string, flags []*cli.Flag, flagValues map[string]interface{}) error {
    print(strings.Repeat("hello", flagValues["count"].(int)))
    return nil
})

cmd.AddFlag(&Flag{
    Name: "count",
    Type: "int",
    Description: "How many times should it print hello",
    Default: 1,
})
```

If flags overlap, the library will autoamtically use flags specified for the command.

### Running options

You can either run the application using `.Execute` or `.Run`. `.Execute` will return a single error object, while `.Run` will exit the application all together in case of error.
