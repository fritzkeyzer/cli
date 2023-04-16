package main

import (
	"fmt"

	"github.com/fritzkeyzer/cli"
)

// create a cli with the following command structure:
// cli
// 		user
// 			list
// 			create
// 		events
// 			list
// 			create
// 			delete

// Nesting is not restricted to 2 levels, you can have as many as you want.
// This example inlines the subcommands for brevity,
// but commands can all be defined as separate vars.

func main() {
	app := cli.App{
		Name:        "cli",
		Description: "Example cli app",
		SubCmds: []cli.Cmd{
			userCmd,
			eventsCmd,
		},
	}

	app.Run()
}

var userCmd = cli.Cmd{
	Name: "user",
	SubCmds: []cli.Cmd{
		{
			Name:        "list",
			Alias:       "ls",
			Description: "List all users",
			Action: func(args map[string]string) {
				fmt.Println("Listing all users")

				for i := 0; i < 10; i++ {
					fmt.Println("User", i, "etc...")
				}
			},
		},
		{
			Name:        "create",
			Description: "Create a new user",
			Args:        []string{"name"},
			Action: func(args map[string]string) {
				fmt.Printf("Creating new user: '%s'\n", args["name"])
			},
		},
	},
}

var eventsCmd = cli.Cmd{
	Name:        "events",
	Description: "Collection of queries to analyse event data",
	SubCmds: []cli.Cmd{
		{
			Name:        "list",
			Alias:       "ls",
			Description: "List all events",
			Action: func(args map[string]string) {
				fmt.Println("Listing all events")
			},
		},

		{
			Name:        "create",
			Description: "List all events",
			Args:        []string{"name"},
			Action: func(args map[string]string) {
				fmt.Printf("Creating event with name='%s'\n", args["name"])
			},
		},

		{
			Name:        "delete",
			Description: "Delete an event",
			Args:        []string{"id"},
			Action: func(args map[string]string) {
				fmt.Printf("Deleting event: '%s'\n", args["id"])
			},
		},
	},
}
