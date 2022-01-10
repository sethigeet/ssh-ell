package commands

import (
	"os"

	"github.com/fatih/color"
)

var Connect = &Command{
	Name:    "connect",
	Desc:    "Connect using the set options",
	HelpMsg: `Write a help message for the connect command here`,

	Execute: func(args ...string) {
		if len(args) != 0 {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid number of args specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "The connect command does not take any args")
			return
		}

		if conn.Connected {
			color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "You are already connected to %q!\n", conn.Host)
			return
		}

		err := conn.Connect()
		if err != nil {
			color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "An error occurred while trying to connect to %q\n", conn.Host)
			color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
			return
		}
	},
}
