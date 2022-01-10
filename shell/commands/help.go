package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"

	"github.com/sethigeet/ssh-ell/utils"
)

var helpMessage = `Write a generic help message for ssh-ell here`

var Help = &Command{
	Name:    "help",
	Desc:    "Print out help info",
	HelpMsg: `Write a help message for the help command here`,

	Complete: func(d prompt.Document) []prompt.Suggest {
		cmd := utils.ParseCmd(d.TextBeforeCursor())
		if len(cmd) > 2 {
			return nil
		}

		return prompt.FilterHasPrefix(
			append(GetCmdCompletionItems(Commands), prompt.Suggest{Text: "help", Description: "Print out help info"}),
			d.GetWordBeforeCursor(),
			true,
		)
	},

	Execute: func(args ...string) {
		if len(args) > 1 {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid args specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "The help command can only take one arg at max!")
			return
		}

		printHeader := func(name, msg string) {
			header := fmt.Sprintf("Help for %s:", name)
			color.New(color.FgGreen, color.Bold).Println(header)
			log := color.New(color.FgGreen)
			log.Println(strings.Repeat("-", len(header)))
			log.Println(msg)
		}

		if len(args) == 0 {
			printHeader("ssh-ell", helpMessage)
			return
		}

		if args[0] == "help" {
			printHeader("help", `Write a help message for the help command here`)
		}

		cmd := GetCmdByName(Commands, args[0])
		if cmd == nil {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid command specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "Please ask for help for a command that exists!")
			return
		}

		printHeader(cmd.Name, cmd.HelpMsg)
	},
}
