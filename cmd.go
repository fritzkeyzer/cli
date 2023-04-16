package cli

import (
	"fmt"
	"strings"
)

type Cmd struct {
	Name        string
	Alias       string   // can be used instead of Name (ideally 1 or 2 characters)
	Description string   // used for documentation
	SubCmds     []Cmd    // sub commands list
	ReqFlags    []Flag   // required flags: if not provided, the cli will print an error, the help doc, and exit
	OptFlags    []Flag   // optional flags: if not provided, the default value will be used. Note that the help flag is automatically added to this list.
	Args        []string // positional args: used to populate the args map passed to the action function.
	Action      func(args map[string]string)

	fullPath []string // for internal use only, to keep track of the full path to the command
}

func (cmd *Cmd) run(args []string, cmdPath []string) error {
	cmd.fullPath = cmdPath

	// check if args contain a sub command
	if len(args) > 0 {
		subCmdName := args[0]

		for _, subCmd := range cmd.SubCmds {
			if subCmd.matchName(subCmdName) {
				return subCmd.run(args[1:], append(cmdPath, subCmdName))
			}
		}
	}

	// this command has no action, print help
	if len(args) == 0 && cmd.Action == nil {
		cmd.printHelp()
		return nil
	}

	// always check for the help flag first
	for _, arg := range args {
		if arg == "-h" || arg == "-help" {
			cmd.printHelp()
			return nil
		}
	}

	// extract flags from args
	var flagArgs []string
	args, flagArgs = splitFlagArgs(args)

	// load required flags
	// if any required flags are not provided, print help and exit
	var flagErrs []error
	for _, fl := range cmd.ReqFlags {
		found, val, err := LoadFlagFromArgs(fl.GetName(), fl.GetAlias(), flagArgs)
		if err != nil {
			flagErrs = append(flagErrs, fmt.Errorf("load flag from args: '%s': %w", fl.GetName(), err))
			continue
		}

		loaded, err := fl.Load(found, val)
		if err != nil {
			flagErrs = append(flagErrs, fmt.Errorf("loading flag: '%s': %w", fl.GetName(), err))
			continue
		}

		if !loaded {
			flagErrs = append(flagErrs, fmt.Errorf("flag: '%s' not provided", fl.GetName()))
		}
	}

	// load optional flags
	for _, fl := range cmd.OptFlags {
		found, val, err := LoadFlagFromArgs(fl.GetName(), fl.GetAlias(), flagArgs)
		if err != nil {
			flagErrs = append(flagErrs, fmt.Errorf("load flag from args: '%s': %w", fl.GetName(), err))
			continue
		}

		_, err = fl.Load(found, val)
		if err != nil {
			flagErrs = append(flagErrs, fmt.Errorf("loading flag: '%s': %w", fl.GetName(), err))
			continue
		}
	}

	// print help and exit if any errors were encountered loading flags
	if len(flagErrs) > 0 {
		cmd.printHelp()

		flagErrStr := "cmd flag error: "
		if len(flagErrs) > 1 {
			flagErrStr = "cmd flag errors: \n\t"
		}
		for i, e := range flagErrs {
			if i > 0 {
				flagErrStr += "\n\t"
			}

			flagErrStr += e.Error()
		}

		return fmt.Errorf(flagErrStr)
	}

	// map out positional args
	argsMap := make(map[string]string)
	for i, argId := range cmd.Args {
		if i >= len(args) {
			break
		}

		argsMap[argId] = args[i]
	}

	// run action
	cmd.Action(argsMap)

	return nil
}

// matchName checks if the given name matches the command name or alias (case-insensitive)
func (cmd *Cmd) matchName(name string) bool {
	name = strings.ToLower(name)

	if name == strings.ToLower(cmd.Name) {
		return true
	}

	if cmd.Alias != "" && name == strings.ToLower(cmd.Alias) {
		return true
	}

	return false
}

// splitFlagArgs from the command args and positional args
func splitFlagArgs(args []string) (remainingArgs []string, flags []string) {
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)
		} else {
			remainingArgs = append(remainingArgs, arg)
		}
	}

	return remainingArgs, flags
}
