// Package commands provides the completion and the execution logic for the
// commands of the shell package
package commands

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"

	"github.com/sethigeet/ssh-ell/ssh"
	"github.com/sethigeet/ssh-ell/utils"
)

var conn = &ssh.Connection{}

func init() {
	conn.ApplyDefaults()
}

type Command struct {
	Name, Desc, HelpMsg string

	Complete func(prompt.Document) []prompt.Suggest
	Execute  func(...string)
}

func (cmd *Command) GetSuggestion() prompt.Suggest {
	return prompt.Suggest{Text: cmd.Name, Description: cmd.Desc}
}

// NOTE: The Commands array does not contain the help command as that ends up
// creating an initialization cycle and hence the help command is accounted for
// in Complete and Execute functions below

// Commands is the list of the commands supported by the shell package
var Commands = []*Command{Connect, Disconnect, Download, Get, Set, Shell, Status, Upload}

func Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return prompt.FilterHasPrefix(
			append(GetCmdCompletionItems(Commands), Help.GetSuggestion()),
			d.GetWordBeforeCursor(),
			true,
		)
	}

	cmd := GetCmdByName(append(Commands, Help), utils.ParseCmd(d.TextBeforeCursor())[0])
	if cmd != nil {
		if cmd.Complete != nil {
			return cmd.Complete(d)
		}
	}

	return prompt.FilterHasPrefix(
		append(GetCmdCompletionItems(Commands), Help.GetSuggestion()),
		d.GetWordBeforeCursor(),
		true,
	)
}

func Execute(s string) {
	s = strings.TrimSpace(s)
	switch s {
	case "":
		return
	case "quit", "exit":
		color.New(color.FgBlue, color.Bold).Println("Byee!")
		os.Exit(0)
		return
	}

	cmd := utils.ParseCmd(s)
	c := GetCmdByName(append(Commands, Help), cmd[0])
	if c != nil {
		c.Execute(cmd[1:]...)
		return
	}

	// Invalid command entered
	color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid command specified")
	color.New(color.FgRed).Fprintln(os.Stderr, "Run `help` to see what commands you can run!")
}
