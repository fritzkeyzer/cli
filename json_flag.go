package cli

import (
	"encoding/json"
	"fmt"
	"os"
)

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

func (flag *JSONFlag[T]) GetName() string {
	return flag.Name
}

func (flag *JSONFlag[T]) GetAlias() string {
	return flag.Alias
}

func (flag *JSONFlag[T]) GetDescription() string {
	desc := flag.Description

	if flag.EnvVar != "" {
		desc += fmt.Sprintf("\n> env var: %s", flag.EnvVar)
	}
	
	return desc
}

func (flag *JSONFlag[T]) Load(argFound bool, argVal *string) (loaded bool, err error) {
	if argFound && argVal == nil {
		return true, fmt.Errorf("no value found")
	}

	if argFound && argVal != nil {
		if err := json.Unmarshal([]byte(*argVal), &flag.Value); err != nil {
			return false, fmt.Errorf("invalid json: %w, loaded from args: '%s'", err, *argVal)
		}

		return true, nil
	}

	envVal := os.Getenv(flag.EnvVar)
	if envVal != "" {
		if err := json.Unmarshal([]byte(envVal), &flag.Value); err != nil {
			return false, fmt.Errorf("invalid json: %w, loaded from env: '%s'", err, envVal)
		}

		return true, nil
	}

	return false, nil
}
