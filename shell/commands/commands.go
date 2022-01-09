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
	{Text: "connect", Description: "Connect using the set options"},
	{Text: "disconnect", Description: "Disconnect from the ssh connection"},
	{Text: "download", Description: "Download a file from the connected server"},
	{Text: "get", Description: "Get the value of an option"},
	{Text: "help", Description: "Print out help info"},
	{Text: "set", Description: "Set the value of an option"},
	{Text: "status", Description: "Get the status of the connection"},
	{Text: "upload", Description: "Upload a file to the connected server"},
}

// Options is the list of options supported by the `set` and `get` commands
// which set and get the options required by other commands
var Options = []prompt.Suggest{
	{Text: "authMethod", Description: "The auth method you want to use to connect to the host"},
	{Text: "host", Description: "The host you want to connect to"},
	{Text: "port", Description: "The port of the host to which you want to connect to"},
	{Text: "timeout", Description: "The timeout for the ssh connection"},
	{Text: "user", Description: "The user on the host to which you want to connect to"},
}

var conn = &ssh.Connection{}

func init() {
	conn.ApplyDefaults()
}

func Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return prompt.FilterHasPrefix(Commands, d.GetWordBeforeCursor(), true)
	}

	switch strings.Split(d.TextBeforeCursor(), " ")[0] {
	case "connect":
		return (Connect{}).Complete(d)
	case "disconnect":
		return (Disconnect{}).Complete(d)
	case "download":
		return (Download{}).Complete(d)
	case "get":
		return (Get{}).Complete(d)
	case "help":
		return (Help{}).Complete(d)
	case "set":
		return (Set{}).Complete(d)
	case "status":
		return (Status{}).Complete(d)
	case "upload":
		return (Upload{}).Complete(d)
	}

	return prompt.FilterHasPrefix(Commands, d.GetWordBeforeCursor(), true)
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

	cmd := strings.Split(s, " ")
	switch cmd[0] {
	case "connect":
		(Connect{}).Execute(cmd[1:]...)
	case "disconnect":
		(Disconnect{}).Execute(cmd[1:]...)
	case "download":
		(Download{}).Execute(cmd[1:]...)
	case "get":
		(Get{}).Execute(cmd[1:]...)
	case "help":
		(Help{}).Execute(cmd[1:]...)
	case "set":
		(Set{}).Execute(cmd[1:]...)
	case "status":
		(Status{}).Execute(cmd[1:]...)
	case "upload":
		(Upload{}).Execute(cmd[1:]...)
	default:
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid command specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "Run `help` to see what commands you can run!")
	}
}
