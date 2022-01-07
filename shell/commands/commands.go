// Package commands provides the completion and the execution logic for the
// commands of the shell package
package commands

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/sethigeet/ssh-ell/ssh"
)

// Commands is the list of the commands supported by the shell package
var Commands = []prompt.Suggest{
	{Text: "help", Description: "Print out help info"},
	{Text: "set", Description: "Set the value of an option"},
	{Text: "get", Description: "Get the value of an option"},
	{Text: "connect", Description: "Connect using the set options"},
	{Text: "upload", Description: "Upload a file to the connected server"},
	{Text: "download", Description: "Download a file from the connected server"},
}

// Options is the list of options supported by the `set` and `get` commands
// which set and get the options required by other commands
var Options = []prompt.Suggest{
	{Text: "host", Description: "The host you want to connect to"},
	{Text: "port", Description: "The port of the host to which you want to connect to"},
	{Text: "user", Description: "The user on the host to which you want to connect to"},
	{Text: "timeout", Description: "The timeout for the ssh connection"},
	{Text: "authMethod", Description: "The auth method you want to use to connect to the host"},
}

var conn ssh.Connection

func Complete(d prompt.Document) []prompt.Suggest {
	switch strings.Split(d.TextBeforeCursor(), " ")[0] {
	case "help":
		return (Help{}).Complete(d)
	case "set":
		return (Set{}).Complete(d)
	case "get":
		return (Get{}).Complete(d)
	case "connect":
		return (Connect{}).Complete(d)
	case "upload":
		return (Upload{}).Complete(d)
	case "download":
		return (Download{}).Complete(d)
	}

	return prompt.FilterFuzzy(Commands, d.GetWordBeforeCursor(), true)
}

func Execute(s string) {
	cmd := strings.Split(s, " ")
	switch cmd[0] {
	case "help":
		(Help{}).Execute(cmd[1:]...)
	case "set":
		(Help{}).Execute(cmd[1:]...)
	case "get":
		(Help{}).Execute(cmd[1:]...)
	case "connect":
		(Help{}).Execute(cmd[1:]...)
	case "upload":
		(Help{}).Execute(cmd[1:]...)
	case "download":
		(Help{}).Execute(cmd[1:]...)
	default:
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid command specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "Run `help` to see what commands you can run!")
	}
}
