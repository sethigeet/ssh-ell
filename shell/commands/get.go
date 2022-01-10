package commands

import (
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"

	"github.com/sethigeet/ssh-ell/shell/commands/options"
	"github.com/sethigeet/ssh-ell/utils"
)

var Get = &Command{
	Name:    "get",
	Desc:    "Get the value of an option",
	HelpMsg: `Write a help message for the get command here`,

	Complete: func(d prompt.Document) []prompt.Suggest {
		cmd := utils.ParseCmd(d.TextBeforeCursor())
		if len(cmd) > 2 {
			return nil
		}

		return prompt.FilterHasPrefix(options.GetOptCompletionItems(), d.GetWordBeforeCursor(), true)
	},

	Execute: func(args ...string) {
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
			for _, opt := range options.Options {
				table.Append([]string{opt.Name, opt.Get(conn)})
			}
			table.Render()
			return
		}
		opt := options.GetOptByName(args[0])
		if opt == nil {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid option specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "Run `help set` to see the list of available options!")
			return
		}

		table.Append([]string{opt.Name, opt.Get(conn)})

		table.Render()
	},
}
