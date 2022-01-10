package commands

import (
	"os"

	"github.com/fatih/color"
)

var Shell = &Command{
	Name:    "shell",
	Desc:    "Start an interactive shell on the connected host",
	HelpMsg: `Write a help message for the shell command here`,

	Execute: func(args ...string) {
		if len(args) != 0 {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid number of args specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "The shell command does not take any args")
			return
		}

		if !conn.Connected {
			color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "You are not connected to any host!\n")
			return
		}

		err := conn.Shell()
		if err != nil {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "An error occurred while trying to start a shell")
			color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
			return
		}
	},
}
