package cli

import (
	"fmt"
	"strings"
)

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
	// Should return true if the flag was loaded  (used for required/optional validation)
	// Returns an error if the flag was found but the value was invalid etc.
	// This method can be used to load the flag from any source, not just cli args. Eg: env vars.
	Load(argFound bool, argVal *string) (loaded bool, err error)
}

// LoadFlagFromArgs will load a flag from cli args.
// Returns true if the flag was found, false otherwise.
// If the flag was found, value will contain the value provided for the flag, if any.
// If the flag was found with a '=' and no value after the '=' an error will be returned.
func LoadFlagFromArgs(name, alias string, args []string) (found bool, value *string, err error) {
	if name != "" {
		found, value, err = loadFlagFromArgsFormatted(formatFlag(name), args)
		if found {
			return true, value, err
		}
	}

	if alias != "" && !found {
		found, value, err = loadFlagFromArgsFormatted(formatAlias(alias), args)
		if found {
			return true, value, err
		}
	}

	return false, nil, nil
}

func formatFlag(name string) string {
	return "--" + name
}

func formatAlias(name string) string {
	return "-" + name
}

// loadFlagFromArgsFormatted will load a flag from cli args.
// flag must be the formatted version of the name or alias. Eg: --flag or -f
func loadFlagFromArgsFormatted(flag string, args []string) (found bool, value *string, err error) {
	for _, arg := range args {
		// flags without values:
		if arg == flag {
			return true, nil, nil
		}

		if !strings.HasPrefix(arg, flag+"=") {
			continue
		}

		// check all formats with values:
		// 	flag="value"
		// 	flag='value'
		// 	flag=value

		// trim front
		trimmed := strings.TrimPrefix(arg, flag+"=")

		if len(trimmed) == 0 {
			return true, nil, fmt.Errorf("flag '%s' found: '%s' but no value provied", flag, arg)
		}

		if strings.HasPrefix(trimmed, "\"") && strings.HasSuffix(trimmed, "\"") {
			trimmed = strings.TrimPrefix(trimmed, "\"")
			trimmed = strings.TrimSuffix(trimmed, "\"")

			return true, &trimmed, nil
		}

		if strings.HasPrefix(trimmed, "'") && strings.HasSuffix(trimmed, "'") {
			trimmed = strings.TrimPrefix(trimmed, "'")
			trimmed = strings.TrimSuffix(trimmed, "'")

			return true, &trimmed, nil
		}

		return true, &trimmed, nil
	}

	return false, nil, nil
}
