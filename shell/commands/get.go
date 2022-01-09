package commands

import (
	"os"
	"strconv"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"

	"github.com/sethigeet/ssh-ell/utils"
)

type Get struct{}

func (Get) Complete(d prompt.Document) []prompt.Suggest {
	cmd := utils.ParseCmd(d.TextBeforeCursor())
	if len(cmd) > 2 {
		return nil
	}

	return prompt.FilterHasPrefix(Options, d.GetWordBeforeCursor(), true)
}

func (Get) Execute(args ...string) {
	if len(args) > 1 {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid number of args specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "The get command take only one or zero args")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Option", "Value"})
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor},
	)

	table.SetAlignment(tablewriter.ALIGN_LEFT)

	if len(args) == 0 {
		table.AppendBulk([][]string{
			{"authMethod", conn.AuthMethodCommonName},
			{"host", conn.Host},
			{"port", strconv.Itoa(int(conn.Port))},
			{"timeout", conn.Timeout.String()},
			{"user", conn.User},
		})
	} else {
		opt := args[0]
		switch args[0] {
		case "authMethod":
			table.Append([]string{opt, conn.AuthMethodCommonName})
		case "host":
			table.Append([]string{opt, conn.Host})
		case "port":
			table.Append([]string{opt, strconv.Itoa(int(conn.Port))})
		case "timeout":
			table.Append([]string{opt, conn.Timeout.String()})
		case "user":
			table.Append([]string{opt, conn.User})
		default:
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid option specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "Run `help set` to see the list of available options!")
			return
		}
	}

	table.Render()
}
