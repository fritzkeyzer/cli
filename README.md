[![Go Report Card](https://goreportcard.com/badge/github.com/fritzkeyzer/cli)](https://goreportcard.com/report/github.com/fritzkeyzer/cli)

# CLI
A go package for creating command line interfaces.

It is slightly opinionated, but it is easy to use and extend.

## Key features
Commands and flags are defined as structs, which makes code easy to read and compose. 

Flags can be specified as required or optional and are loaded and validated before the commands are executed

Various flag types are included: string, int and a json flag type utilising generics.
Additional flag types can be created by implementing the Flag interface.

Sub-commands are supported, allowing for a tree-like structure - at the same time keeping boilerplate code for positional args and flags to a minimum. 

Positional arguments are loaded into a map, allowing access by name. 
This improves readability and removes the need for checking the length of the args slice.

```go
// this is nicer
name := args["name"]

// than this
name := ""
if len(args) > 0 {
    name = args[0]
}
```



## An example cli
```go
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
	Value: 5, // default value
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


```
Running this example:
```bash
$ cli
```
```
An example CLI demonstrating different flag types

Usage:
    cli [command] [flags]

Commands:
    db    database    Do something with a database
          hello       Say hello to <name> a number of times
          person      Add a person to a database or something

Optional Flags:
    -h    --help    Print documentation for command
```

```bash
$ cli db
```
```
Do something with a database

Usage:
    cli db [flags]

Required Flags:
         --db-conn    Database connection string
                      > env var: DB_CONN

Optional Flags:
    -h    --help    Print documentation for command

2023/04/16 16:20:18 ERROR: cmd flag error: flag: 'db-conn' not provided
```
```bash
$ cli hello --help
```
```
Say hello to <name> a number of times

Usage:
    cli hello <name> [flags]

Optional Flags:
    -n    --count
    -h    --help     Print documentation for command
```
```bash
$ cli hello 'freddy the gopher' --count=3
```
```
Hello freddy the gopher 0
Hello freddy the gopher 1
Hello freddy the gopher 2
```
```bash
$ cli person
```
```
Add a person to a database or something

Usage:
    cli person [flags]

Required Flags:
         --person    A person in JSON format

Optional Flags:
    -h    --help    Print documentation for command

2023/04/16 16:20:45 ERROR: cmd flag error: flag: 'person' not provided
```
```bash
$ cli person --person='{"name":"fritz", "age":25}'
```
```
2023/04/16 16:20:48 Person name: fritz
2023/04/16 16:20:48 Person age: 25
```


## Flags
### Included in this package:
```go
// StringFlag is a flag that which can be provided by cli args or env var.
// CLI args take precedence.
// If AcceptedValues are specified, the value is validated against them.
type StringFlag struct {
    Name           string
    Alias          string
    EnvVar         string
    Description    string
    AcceptedValues []string // if specified, only these values are accepted
    Value          string   // can provide a default value here
}

// IntFlag can be provided by cli args or env var. Values are parsed as int.
// CLI args take precedence.
type IntFlag struct {
    Name        string
    Alias       string
    Description string
    Value       int // can provide a default value here
}

// JSONFlag will parse a json string into Value, of type T.
// The value can be provided by cli args or env var.
// CLI args take precedence.
type JSONFlag[T any] struct {
    Name        string
    Alias       string
    EnvVar      string
    Description string
    Value       T // can provide a default value here
}
```

### Flag interface
```go
// Flag allows for custom flag types to be created.
// (Although the included flag types should be sufficient for most cases.)
// Command line flags are expected in one of the following formats:
//
//	--flag=value
//	--flag='value'
//	--flag="value"
//	--flag
//
// Or using an alias:
//
//	-f=value
//	-f='value'
//	-f="value"
//	-f
type Flag interface {
    GetName() string
    GetAlias() string
    GetDescription() string
    
    // Load a flag. The argFound bool indicates if the flag was found in the cli args
    // and argVal contains the value provided for the flag, if any.
    // Should return true if the flag was loaded (used for required/optional validation)
    // Returns an error if the flag was found but the value was invalid etc.
    // This method can be used to load the flag from any source, not just cli args. Eg: env vars.
    Load(argFound bool, argVal *string) (loaded bool, err error)
}
```
