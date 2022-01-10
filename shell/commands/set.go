package commands

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"

	"github.com/sethigeet/ssh-ell/shell/commands/options"
	"github.com/sethigeet/ssh-ell/utils"
)

var Set = &Command{
	Name:    "set",
	Desc:    "Set the value of an option",
	HelpMsg: `Write a help message for the set command here`,

	Complete: func(d prompt.Document) []prompt.Suggest {
		cmd := utils.ParseCmd(d.TextBeforeCursor())
		if len(cmd) > 3 {
			return nil
		}

		if len(cmd) == 3 {
			if cmd[1] == "authMethod" {
				return prompt.FilterHasPrefix(options.AuthMethodOptions, d.GetWordBeforeCursor(), true)
			}
			return nil
		}

		return prompt.FilterHasPrefix(options.GetOptCompletionItems(), d.GetWordBeforeCursor(), true)
	},

	Execute: func(args ...string) {
		if len(args) != 2 {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid args specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "The set command take two arguments, the first being the option and the second being the value.")
			return
		}

		if conn.Connected {
			color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "You cannot set options while you are connected to a host!\n")
			return
		}

		opt := options.GetOptByName(args[0])
		if opt == nil {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid option specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "Run `help set` to see the options that you can set!")
			return
		}
		err := opt.Set(strings.Join(args[1:], " "), conn)
		if err != nil {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "An error occurred while setting the option")
			color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
		}
	},
}
