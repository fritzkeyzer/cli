package cli

import (
	"fmt"
	"os"
	"strings"
)

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

func (flag *StringFlag) GetName() string {
	return flag.Name
}

func (flag *StringFlag) GetAlias() string {
	return flag.Alias
}

func (flag *StringFlag) GetDescription() string {
	desc := flag.Description

	if flag.EnvVar != "" {
		if desc != "" {
			desc += "\n"
		}
		desc += "> env var: " + flag.EnvVar
	}

	if len(flag.AcceptedValues) != 0 {
		desc += fmt.Sprintf("\n> accepted values: [%s]", strings.Join(flag.AcceptedValues, ", "))
	}

	return desc
}

func (flag *StringFlag) Load(argFound bool, argVal *string) (loaded bool, err error) {
	if argFound && argVal == nil {
		return true, fmt.Errorf("no value found")
	}

	if argFound && argVal != nil {
		flag.Value = *argVal
		return true, flag.validateVal()
	}

	envVal := os.Getenv(flag.EnvVar)
	if envVal != "" {
		flag.Value = envVal
		return true, flag.validateVal()
	}

	return false, nil
}

func (flag *StringFlag) validateVal() error {
	if len(flag.AcceptedValues) == 0 {
		return nil
	}

	for _, acceptedValue := range flag.AcceptedValues {
		if flag.Value == acceptedValue {
			return nil
		}
	}

	return fmt.Errorf("'%s' is not an accepted value", flag.Value)
}
