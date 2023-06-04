package main

import (
	"fmt"

	"github.com/fritzkeyzer/cli"
)

func main() {
	app := cli.App{
		Name:        "cli",
		Description: "An example CLI",
		SubCmds: []cli.Cmd{
			createUserCmd,
		},
	}

	app.Run()
}

// createUserCmd is a simple cmd using a positional arg
var createUserCmd = cli.Cmd{
	Name:        "create",
	Description: "Create a new user",
	Args:        []string{"name"}, // name is a positional arg
	Action: func(args map[string]string) {
		name := args["name"]
		fmt.Println("Creating user:", name)
	},
}
