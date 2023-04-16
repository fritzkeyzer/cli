package main

import (
	"fmt"
	"log"

	"github.com/fritzkeyzer/cli"
)

var dbConnFlag = &cli.StringFlag{
	Name:        "db-conn",
	EnvVar:      "DB_CONN",
	Description: "Database connection string",
}

var countFlag = &cli.IntFlag{
	Name:  "count",
	Alias: "n",
	Value: 5, // default value, since countFlag is optional
}

// personFlag demonstrates a generic JSONFlag, using the Person type
var personFlag = &cli.JSONFlag[Person]{
	Name:        "person",
	Description: "A person in JSON format",
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	app := cli.App{
		Name:        "cli",
		Description: "An example CLI demonstrating different flag types",
		SubCmds: []cli.Cmd{
			dbCmd,
			helloCmd,
			personCmd,
		},
	}

	app.Run()
}

var dbCmd = cli.Cmd{
	Name:        "database",
	Alias:       "db", // commands can have aliases
	Description: "Do something with a database",
	ReqFlags:    []cli.Flag{dbConnFlag},
	Action: func(args map[string]string) {
		log.Println("Database things")
		log.Println("Connection string:", dbConnFlag.Value)
	},
}

var helloCmd = cli.Cmd{
	Name:        "hello",
	Description: "Say hello to <name> a number of times",
	OptFlags:    []cli.Flag{countFlag},
	Args:        []string{"name"},
	Action: func(args map[string]string) {
		for i := 0; i < countFlag.Value; i++ {
			fmt.Println("Hello", args["name"], i)
		}
	},
}

var personCmd = cli.Cmd{
	Name:        "person",
	Description: "Add a person to a database or something",
	ReqFlags:    []cli.Flag{personFlag},
	Action: func(args map[string]string) {
		person := personFlag.Value

		log.Println("Person name:", person.Name)
		log.Println("Person age:", person.Age)
	},
}
