[![Go Report Card](https://goreportcard.com/badge/github.com/fritzkeyzer/cli)](https://goreportcard.com/report/github.com/fritzkeyzer/cli)

# CLI
CLI is a go package for creating command line interfaces.
It is intended to be very simple to use and easy to read.

## Commands
Commands and flags are declarative, which makes code easy to read and compose.

Positional arguments are loaded into a map, allowing access by name.
This improves readability and developer experience.
Positional args are included in the help text.

```go
func main(){
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
```

## Sub-commands
Every command can have sub-commands, allowing for a tree-like structure.

```go
var userCmd = cli.Cmd{
    Name:        "user",
    Description: "Manage users",
    SubCmds: []cli.Cmd{
        createUserCmd,
        listUsersCmd,
    },
}
```

## Flags
Flags are defined as an interface, allowing for custom flag types to be created.

A number of flag types are included in this package: `StringFlag`, `IntFlag` and `JSONFlag`.

>Note that there are additional options for these flags that have not been set
```go
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
```


## Optional/required flags 

Flags can be specified as required or optional and are loaded and validated before the commands are executed.

```go
var doSomethingCmd = cli.Cmd{
    Name:        "something",
    Description: "Performs some action",
    ReqFlags:    []cli.Flag{dbConnFlag, apiKeyFlag},
    OptFlags:    []cli.Flag{envNameFlag},
    Action: func(args map[string]string) {
        // dbConnFlag and apiKeyFlag are guaranteed to be set
        // envNameFlag is optional, if it was not provided it will use the default value
        
        // do something with the flags...
        fmt.Println("dbConn:", dbConnFlag.Value)
        fmt.Println("apiKey:", apiKeyFlag.Value)
        fmt.Println("env:", envNameFlag.Value)
    },
}
```

## Flag interface
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
    // This method can be used to load the flag from any source, not just cli args. Eg: env vars, files etc.
    Load(argFound bool, argVal *string) (loaded bool, err error)
}
```

## Example cli output

See `_examples` folder for full code examples

`$ cli`
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

`$ cli db`
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

`$ cli hello --help`
```
Say hello to <name> a number of times

Usage:
    cli hello <name> [flags]

Optional Flags:
    -n    --count
    -h    --help     Print documentation for command
```
  
`$ cli hello 'freddy the gopher' --count=3`
```
Hello freddy the gopher 0
Hello freddy the gopher 1
Hello freddy the gopher 2
```

`$ cli person`
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

`$ cli person --person='{"name":"fritz", "age":25}'`
```
2023/04/16 16:20:48 Person name: fritz
2023/04/16 16:20:48 Person age: 25
```
