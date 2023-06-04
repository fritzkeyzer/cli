package cli

// BoolFlag is a flag that if provided, the value will be true.
// If not provided, the value will be false.
type BoolFlag struct {
	Name        string
	Alias       string
	Description string
	Value       bool
}

func (flag *BoolFlag) GetName() string {
	return flag.Name
}

func (flag *BoolFlag) GetAlias() string {
	return flag.Alias
}

func (flag *BoolFlag) GetDescription() string {
	desc := flag.Description

	return desc
}

func (flag *BoolFlag) Load(argFound bool, argVal *string) (loaded bool, err error) {
	flag.Value = argFound

	return true, nil
}
