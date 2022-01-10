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

func Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return prompt.FilterHasPrefix(Commands, d.GetWordBeforeCursor(), true)
	}

	switch utils.ParseCmd(d.TextBeforeCursor())[0] {
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
	case "shell":
		return (Shell{}).Complete(d)
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

	cmd := utils.ParseCmd(s)
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
	case "shell":
		(Shell{}).Execute(cmd[1:]...)
	case "upload":
		(Upload{}).Execute(cmd[1:]...)
	default:
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid command specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "Run `help` to see what commands you can run!")
	}
}
