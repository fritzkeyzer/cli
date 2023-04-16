package cli

import (
	"log"
	"os"
)

type App struct {
	Name        string   // for documentation only (should match the name of the executable)
	Description string   // used for documentation
	SubCmds     []Cmd    // a list of available sub commands
	ReqFlags    []Flag   // required flags: if not provided, the cli will print an error, the help doc, and exit
	OptFlags    []Flag   // optional flags: if not provided, the default value will be used. Note that the help flag is automatically added to this list.
	Args        []string // positional args: used to populate the args map passed to the action function.
	Action      func(args map[string]string)
}

func (app *App) Run() {
	rootCmd := Cmd{
		Name:        app.Name,
		Args:        app.Args,
		Description: app.Description,
		SubCmds:     app.SubCmds,
		ReqFlags:    app.ReqFlags,
		OptFlags:    app.OptFlags,
		Action:      app.Action,
	}

	err := rootCmd.run(os.Args[1:], []string{app.Name})
	if err != nil {
		log.Fatal("ERROR: ", err)
	}
}
