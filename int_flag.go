package cli

import (
	"fmt"
	"strconv"
)

// IntFlag can be provided by cli args or env var. Values are parsed as int.
// CLI args take precedence.
type IntFlag struct {
	Name        string
	Alias       string
	Description string
	Value       int // can provide a default value here
}

func (flag *IntFlag) GetName() string {
	return flag.Name
}

func (flag *IntFlag) GetAlias() string {
	return flag.Alias
}

func (flag *IntFlag) GetDescription() string {
	desc := flag.Description

	return desc
}

func (flag *IntFlag) Load(argFound bool, argVal *string) (loaded bool, err error) {
	if !argFound {
		return false, nil
	}
	if argVal == nil {
		return true, fmt.Errorf("flag %s is missing a value", flag.Name)
	}

	if *argVal != "" {
		parsed, err := strconv.ParseInt(*argVal, 10, 64)
		if err != nil {
			return false, fmt.Errorf("parsing int: %w", err)
		}
		flag.Value = int(parsed)
		return true, nil
	}

	return false, nil
}
