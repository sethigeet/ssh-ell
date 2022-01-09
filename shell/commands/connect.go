package commands

import (
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
)

type Connect struct{}

func (Connect) Complete(d prompt.Document) []prompt.Suggest {
	return nil
}

func (Connect) Execute(args ...string) {
	if len(args) > 0 {
		color.New(color.FgRed, color.Bold).Fprintln(os.Stderr, "Invalid number of args specified")
		color.New(color.FgRed).Fprintln(os.Stderr, "The connect command does not take any args")
		return
	}

	err := conn.Connect()
	if err != nil {
		color.New(color.FgRed, color.Bold).Fprintf(os.Stderr, "An error occurred while trying to connect to %q\n", conn.Host)
		color.New(color.FgRed).Fprintf(os.Stderr, "error: %s\n", err)
		return
	}
}
