package commands

import (
	"os"

	"github.com/fatih/color"
)

var Status = &Command{
	Name:    "status",
	Desc:    "Get the status of the connection",
	HelpMsg: `Write a help message for the status command here`,

	Execute: func(args ...string) {
		if len(args) > 0 {
			color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid number of args specified")
			color.New(color.FgRed).Fprintln(os.Stderr, "The status command does not take any args")
			return
		}

		if conn.Connected {
			color.New(color.FgGreen, color.Bold).Printf("Status: ")
			color.New(color.FgGreen).Println("connected")
		} else {
			color.New(color.FgRed, color.Bold).Printf("Status: ")
			color.New(color.FgRed).Println("not connected")
		}
	},
}
