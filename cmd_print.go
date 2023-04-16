package cli

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

const descriptionWrapLimit = 50

var helpFlag = StringFlag{
	Name:        "help",
	Alias:       "h",
	Description: `Print documentation for command`,
}

func (cmd *Cmd) printHelp() {
	// Description
	if cmd.Description != "" {
		fmt.Println(cmd.Description)
		fmt.Println()
	}

	// Args
	fmt.Println("Usage:")
	usageText := strings.Join(cmd.fullPath, " ")
	if len(cmd.SubCmds) > 0 {
		usageText += " [command]"
	}
	for _, arg := range cmd.Args {
		usageText += fmt.Sprintf(" <%s>", arg)
	}
	usageText += " [flags]"

	fmt.Println("    " + usageText)
	fmt.Println()

	printCommandsSection("Commands:", cmd.SubCmds)

	printFlagsSection("Required Flags:", cmd.ReqFlags)

	printFlagsSection("Optional Flags:", append(cmd.OptFlags, &helpFlag))
}

func printCommandsSection(title string, cmds []Cmd) {
	if len(cmds) == 0 {
		return
	}

	fmt.Println(title)

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	for _, cmd := range cmds {

		alias := " "
		if cmd.Alias != "" {
			alias = cmd.Alias
		}

		name := cmd.Name

		descLines := wrapText(cmd.Description, descriptionWrapLimit)

		for i := range descLines {
			if i == 0 {
				tw.Write([]byte(fmt.Sprintf("\t%s\t%s\t%s\n", alias, name, descLines[i])))
				continue
			}

			tw.Write([]byte(fmt.Sprintf("\t%s\t%s\t%s\n", "", "", descLines[i])))
		}

	}
	tw.Flush()
	fmt.Println()
}

func printFlagsSection(title string, flags []Flag) {
	if len(flags) == 0 {
		return
	}

	fmt.Println(title)

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
	for _, flag := range flags {

		alias := " "
		if flag.GetAlias() != "" {
			alias = formatAlias(flag.GetAlias())
		}

		name := formatFlag(flag.GetName())

		descLines := wrapText(flag.GetDescription(), descriptionWrapLimit)

		for i := range descLines {
			if i == 0 {
				tw.Write([]byte(fmt.Sprintf("\t%s\t%s\t%s\n", alias, name, descLines[i])))
				continue
			}

			tw.Write([]byte(fmt.Sprintf("\t%s\t%s\t%s\n", "", "", descLines[i])))
		}

	}
	tw.Flush()
	fmt.Println()
}

// wrapText takes multiline text and re-wraps it to ensure it fits with the specified limit
func wrapText(text string, maxWidth int) []string {
	// for safety
	if maxWidth < 1 {
		maxWidth = 1
	}

	inputLines := strings.Split(text, "\n")
	var outputLines []string
	for i := range inputLines {
		if len(inputLines[i]) <= maxWidth {
			outputLines = append(outputLines, inputLines[i])
			continue
		}

		// chop line
		chopIndex := strings.LastIndex(inputLines[i][:maxWidth], " ")
		if chopIndex == -1 {
			chopIndex = maxWidth
		}

		if chopIndex == 0 {
			outputLines = append(outputLines, inputLines[i])
			continue
		}

		if chopIndex > len(inputLines[i]) {
			chopIndex = len(inputLines[i])
		}

		outputLines = append(outputLines, inputLines[i][:chopIndex])
		remaining := inputLines[i][chopIndex:]
		remaining = strings.TrimPrefix(remaining, " ")

		if remaining == "" {
			continue
		}

		otherLines := wrapText(remaining, maxWidth)
		outputLines = append(outputLines, otherLines...)
	}

	return outputLines
}
