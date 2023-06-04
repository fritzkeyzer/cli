package main

import (
	"fmt"
	"log"

	"github.com/fritzkeyzer/cli"
)

// create a cli with the following command structure:
// cli
// 	user
// 		create
// 		list
// 	events
// 		create
// 		list
// 		delete

// Nesting is not restricted to 2 levels, you can have as many as you want.

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
	Name:        "user",
	Description: "Manage users",
	SubCmds: []cli.Cmd{
		createUserCmd,
		listUsersCmd,
	},
}

var createUserCmd = cli.Cmd{
	Name:        "create",
	Description: "Create a new user",
	Args:        []string{"name"},
	Action: func(args map[string]string) {
		name := args["name"]

		fmt.Println("Creating user:", name)
	},
}

var listUsersCmd = cli.Cmd{
	Name:        "list",
	Alias:       "ls",
	Description: "List all users",
	Action: func(args map[string]string) {
		fmt.Println("Listing users...")

		for i := 0; i < 10; i++ {
			fmt.Println("User", i, "etc...")
		}
	},
}

var eventsCmd = cli.Cmd{
	Name:        "events",
	Description: "Manage events",
	SubCmds: []cli.Cmd{
		createEventCmd,
		listEventsCmd,
		deleteEventCmd,
	},
}

var createEventCmd = cli.Cmd{
	Name:        "create",
	Description: "Create an event",
	Args:        []string{"name"},
	Action: func(args map[string]string) {
		eventName := args["name"]

		fmt.Println("Creating event:", eventName)
	},
}

var listEventsCmd = cli.Cmd{
	Name:        "list",
	Alias:       "ls",
	Description: "List all events",
	Action: func(args map[string]string) {
		fmt.Println("Listing all events")
	},
}

var deleteEventCmd = cli.Cmd{
	Name:        "delete",
	Description: "Delete an event",
	Args:        []string{"id"},
	Action: func(args map[string]string) {
		eventId := args["id"]

		fmt.Println("Deleting event:", eventId)
	},
}

var backupCmd = cli.Cmd{
	Name:        "hello",
	Alias:       "",
	Description: "Create a database backup and upload it to GCS",
	SubCmds:     []cli.Cmd{},  // sub-commands can be nested as deep as you want
	ReqFlags:    []cli.Flag{}, // required flags will prevent execution if not provided
	OptFlags:    []cli.Flag{}, // optional flags will use default values if they are not provided
	Args:        []string{},   // if specified, positional args are loaded into a map, allowing access by name
	Action: func(args map[string]string) {
		log.Println("Perform database backup")

		// database.BackupToGCS(context.Background(),
		// 	DBConnFlag.Value,
		// 	GCPCredsFlag.Value,
		// 	BackupBucketFlag.Value,
		// )
	},
}
